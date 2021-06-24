package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// NaxsiFmtItem include list of matched in naxsi fmt log
type NaxsiFmtItem struct {
	Zone    string `json:"zone"`
	ID      string `json:"id"`
	VarName string `json:"var_name"`
	CScore  string `json:"cscore"`
	Score   string `json:"score"`
}

// NginxErrorEntry is nginx error log packet
type NginxErrorEntry struct {
	// essential
	Time    string `json:"time"`
	Level   string `json:"level"`
	PID     int    `json:"pid"`
	TID     int    `json:"tid"`
	CID     int    `json:"cid"`
	Message string `json:"message"`

	// additional
	Msg string `json:"msg"`

	checkSumParts []string

	checkSumUseMsg     bool
	CheckSum           string `json:"checksum"`
	CheckSumDebug      string `json:"checksum_debug"`
	Client             string `json:"client"`
	Server             string `json:"server"`
	Host               string `json:"host"`
	Upstream           string `json:"upstream"`
	UpstreamHost       string `json:"upstream_host"`
	Referrer           string `json:"referrer"`
	ReferrerHost       string `json:"referrer_host"`
	RequestMethod      string `json:"request_method"`
	RequestURI         string `json:"request_uri"`
	RequestHTTPVersion string `json:"request_http_version"`

	ErrorType    string `json:"error_type"`
	ErrorDetails string `json:"error_details"`

	// naxsi
	NaxsiMode string `json:"naxsi"`

	// naxsi fmt
	NaxsiFmtIP             string         `json:"naxsi_fmt_ip"`
	NaxsiFmtServer         string         `json:"naxsi_fmt_server"`
	NaxsiFmtURI            string         `json:"naxsi_fmt_uri"`
	NaxsiFmtLearning       bool           `json:"naxsi_fmt_learning"`
	NaxsiFmtVers           string         `json:"naxsi_fmt_vers"`
	NaxsiFmtBlock          bool           `json:"naxsi_fmt_block"`
	NaxsiFmtTotalProcessed int            `json:"naxsi_fmt_total_processed"`
	NaxsiFmtTotalBlocked   int            `json:"naxsi_fmt_total_blocked"`
	NaxsiFmtItems          []NaxsiFmtItem `json:"naxsi_fmt_items"`

	// naxsi exlog
	NaxsiExLogIP      string `json:"naxsi_exlog_ip"`
	NaxsiExLogServer  string `json:"naxsi_exlog_server"`
	NaxsiExLogURI     string `json:"naxsi_exlog_uri"`
	NaxsiExLogID      string `json:"naxsi_exlog_id"`
	NaxsiExLogZone    string `json:"naxsi_exlog_zone"`
	NaxsiExLogVarName string `json:"naxsi_exlog_var_name"`
	NaxsiExLogContent string `json:"naxsi_exlog_content"`
}

var entryRegex, _ = regexp.Compile(`^(?s)(?P<time>[0-9]+\/[0-9]+\/[0-9]+ [0-9]+:[0-9]+:[0-9]+) \[(?P<level>[a-z]+)\] (?P<pid>[0-9]+)#(?P<tid>[0-9]+): \*(?P<cid>[0-9]+)(.*)$`)

var nginxErrorLogDateTimeLayout = "2006/01/02 15:04:05"

var noneAlphaRegex, _ = regexp.Compile(`[^a-zA-Z]`)
var oneSpaceRegex, _ = regexp.Compile(`[\s]+`)

func parserTime(timeString string) (unixtimestamp string, e error) {
	t, err := time.Parse(nginxErrorLogDateTimeLayout, timeString)
	if err != nil {
		return "", err
	}
	return t.Format(time.RFC1123Z), nil
}

func stripLog(message string) string {
	s := noneAlphaRegex.ReplaceAllString(message, " ")
	s = oneSpaceRegex.ReplaceAllString(s, " ")
	return strings.ToLower(strings.TrimSpace(s))
}

func replaceMatched(message string, whole string) string {
	return strings.Replace(message, whole, "", -1)
}

// Parser give single line nginx error log and parse it
func Parser(message string) (entry NginxErrorEntry, e error) {
	if isMatched := entryRegex.MatchString(message); isMatched {
		entry = NginxErrorEntry{}
		matched := entryRegex.FindStringSubmatch(message)

		entry.Time, _ = parserTime(matched[1])
		entry.Level = strings.ToLower(strings.TrimSpace(matched[2]))
		entry.PID, _ = strconv.Atoi(matched[3])
		entry.TID, _ = strconv.Atoi(matched[4])
		entry.CID, _ = strconv.Atoi(matched[5])
		entry.Message = strings.TrimSpace(matched[6])
		entry.Msg = entry.Message
		entry.checkSumUseMsg = true

		entry.NaxsiFmtItems = make([]NaxsiFmtItem, 0)

		// general
		findClient(&entry)
		findServer(&entry)
		findHost(&entry)
		findUpstream(&entry)
		findReferrer(&entry)
		findRequest(&entry)

		entry.checkSumParts = []string{entry.Msg}

		// specials
		findErrorFastCGI(&entry)
		findErrorOpenFailed(&entry)

		// naxsi
		findNaxsiFmt(&entry)
		findNaxsiExLog(&entry)

		entry.Msg = stripLog(entry.Msg)
		if entry.checkSumUseMsg {
			entry.checkSumParts = []string{entry.Msg}
		}

		checkSumParts := strings.Join(entry.checkSumParts, ":")

		entry.CheckSumDebug = checkSumParts

		hash := md5.Sum([]byte(checkSumParts))
		entry.CheckSum = hex.EncodeToString(hash[:])

		return entry, nil
	}
	e = errors.New("Packet is not valid")
	return NginxErrorEntry{}, e
}

// IsJSON just check is valid json or not
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// ParserJSON json output from entry log
func ParserJSON(entry NginxErrorEntry) ([]byte, error) {
	return json.Marshal(entry)
}
