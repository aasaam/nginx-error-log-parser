package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	errorTypeUnknown    string = "_"
	errorTypeNaxsiFmt   string = "naxsi_fmt"
	errorTypeNaxsiExLog string = "naxsi_exlog"
	errorTypeFastcgi    string = "fastcgi"
	errorTypeOpenFailed string = "open_failed"
)

type naxsiFmtItem struct {
	Zone    string `json:"zone"`
	ID      string `json:"id"`
	VarName string `json:"var_name"`
	CScore  string `json:"cscore"`
	Score   string `json:"score"`
}

type nginxErrorEntry struct {
	// essential
	Time    string `json:"time"`
	Level   string `json:"level"`
	PID     int    `json:"pid"`
	TID     int    `json:"tid"`
	CID     int    `json:"cid,omitempty"`
	Message string `json:"message"`

	// additional
	Msg *string `json:"msg"`

	checkSumParts []string

	checkSumUseMsg     bool
	CheckSum           string  `json:"checksum"`
	CheckSumDebug      string  `json:"checksum_debug"`
	Client             *string `json:"client,omitempty"`
	Server             *string `json:"server,omitempty"`
	Host               *string `json:"host,omitempty"`
	Upstream           *string `json:"upstream,omitempty"`
	UpstreamHost       *string `json:"upstream_host,omitempty"`
	Referrer           *string `json:"referrer,omitempty"`
	ReferrerHost       *string `json:"referrer_host,omitempty"`
	RequestMethod      *string `json:"request_method,omitempty"`
	RequestURI         *string `json:"request_uri,omitempty"`
	RequestHTTPVersion *string `json:"request_http_version,omitempty"`

	ErrorType    string  `json:"error_type,omitempty"`
	ErrorDetails *string `json:"error_details,omitempty"`

	// naxsi fmt
	NaxsiFmtIP             *string        `json:"naxsi_fmt_ip,omitempty"`
	NaxsiFmtServer         *string        `json:"naxsi_fmt_server,omitempty"`
	NaxsiFmtURI            *string        `json:"naxsi_fmt_uri,omitempty"`
	NaxsiFmtLearning       *bool          `json:"naxsi_fmt_learning,omitempty"`
	NaxsiFmtVers           *string        `json:"naxsi_fmt_vers,omitempty"`
	NaxsiFmtBlock          *bool          `json:"naxsi_fmt_block,omitempty"`
	NaxsiFmtTotalProcessed *int           `json:"naxsi_fmt_total_processed,omitempty"`
	NaxsiFmtTotalBlocked   *int           `json:"naxsi_fmt_total_blocked,omitempty"`
	NaxsiFmtItems          []naxsiFmtItem `json:"naxsi_fmt_items,omitempty"`

	// naxsi exlog
	NaxsiExLogIP      *string `json:"naxsi_exlog_ip,omitempty"`
	NaxsiExLogServer  *string `json:"naxsi_exlog_server,omitempty"`
	NaxsiExLogURI     *string `json:"naxsi_exlog_uri,omitempty"`
	NaxsiExLogID      *string `json:"naxsi_exlog_id,omitempty"`
	NaxsiExLogZone    *string `json:"naxsi_exlog_zone,omitempty"`
	NaxsiExLogVarName *string `json:"naxsi_exlog_var_name,omitempty"`
	NaxsiExLogContent *string `json:"naxsi_exlog_content,omitempty"`
}

var entryCIDRegex = regexp.MustCompile(`^(?s)(?P<time>[0-9]+\/[0-9]+\/[0-9]+ [0-9]+:[0-9]+:[0-9]+) \[(?P<level>[a-z]+)\] (?P<pid>[0-9]+)#(?P<tid>[0-9]+): \*(?P<cid>[0-9]+)(?P<message>.*)$`)
var entryNoCIDRegex = regexp.MustCompile(`^(?s)(?P<time>[0-9]+\/[0-9]+\/[0-9]+ [0-9]+:[0-9]+:[0-9]+) \[(?P<level>[a-z]+)\] (?P<pid>[0-9]+)#(?P<tid>[0-9]+): (?P<message>.*)$`)

var nginxErrorLogDateTimeLayout = "2006/01/02 15:04:05"

var noneAlphaRegex = regexp.MustCompile(`[^a-zA-Z]`)
var oneSpaceRegex = regexp.MustCompile(`[\s]+`)

func parserTime(timeString string) (t time.Time, e error) {
	t, err := time.Parse(nginxErrorLogDateTimeLayout, timeString)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func stripLog(message string) string {
	s := noneAlphaRegex.ReplaceAllString(message, " ")
	s = oneSpaceRegex.ReplaceAllString(s, " ")
	return strings.ToLower(strings.TrimSpace(s))
}

func replaceMatched(message string, whole string) string {
	return strings.Replace(message, whole, "", -1)
}

func parser(message string) (entry nginxErrorEntry, e error) {
	var matched []string
	var re *regexp.Regexp

	if isMatched := entryCIDRegex.MatchString(message); isMatched {
		re = entryCIDRegex
		matched = entryCIDRegex.FindStringSubmatch(message)
	} else if isMatched := entryNoCIDRegex.MatchString(message); isMatched {
		re = entryNoCIDRegex
		matched = entryNoCIDRegex.FindStringSubmatch(message)
	} else {
		return nginxErrorEntry{}, errors.New("packet is not valid")
	}

	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matched[i]
		}
	}

	logTime, logTimeErr := parserTime(result["time"])
	if logTimeErr != nil {
		logTime = time.Now()
	}

	entry = nginxErrorEntry{
		ErrorType:      errorTypeUnknown,
		Time:           logTime.Format(time.RFC1123Z),
		PID:            prettySureInt(result["pid"]),
		TID:            prettySureInt(result["tid"]),
		Level:          strings.ToLower(strings.TrimSpace(result["level"])),
		Message:        strings.TrimSpace(result["message"]),
		checkSumUseMsg: true,
	}

	if result["cid"] != "" {
		entry.CID = prettySureInt(result["cid"])
	}
	entry.Msg = stringPointer(entry.Message)

	entry.NaxsiFmtItems = make([]naxsiFmtItem, 0)

	// general
	findClient(&entry)
	findServer(&entry)
	findHost(&entry)
	findUpstream(&entry)
	findReferrer(&entry)
	findRequest(&entry)

	entry.checkSumParts = []string{*entry.Msg, entry.Level}

	// specials
	findErrorFastCGI(&entry)
	findErrorOpenFailed(&entry)

	// naxsi
	findNaxsiFmt(&entry)
	findNaxsiExLog(&entry)

	entry.Msg = stringPointer(stripLog(*entry.Msg))
	if entry.checkSumUseMsg {
		entry.checkSumParts = []string{*entry.Msg}
	}

	checkSumParts := strings.Join(entry.checkSumParts, ":")

	entry.CheckSumDebug = checkSumParts

	hash := sha1.Sum([]byte(checkSumParts))
	entry.CheckSum = hex.EncodeToString(hash[:])

	return entry, nil
}

func isJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func parserJSON(entry nginxErrorEntry) ([]byte, error) {
	return json.Marshal(entry)
}

func prettySureInt(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return 0
	}
	return i
}

func stringPointer(s string) *string {
	if s == "" {
		return nil
	}
	var pt *string
	pt = &s
	return pt
}

func boolPointer(b bool) *bool {
	var pt *bool
	pt = &b
	return pt
}

func intPointer(i int) *int {
	var pt *int
	pt = &i
	return pt
}
