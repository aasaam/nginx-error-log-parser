package main

import (
	"net/url"
	"regexp"
)

var naxsiExLogRegex, _ = regexp.Compile(`NAXSI_EXLOG: (?P<naxsiExLog>[^ ]+),`)

func findNaxsiExLog(entry *NginxErrorEntry) {
	if ok := naxsiExLogRegex.MatchString(entry.Message); ok {
		matched := naxsiExLogRegex.FindStringSubmatch(entry.Message)
		entry.Msg = replaceMatched(entry.Msg, matched[0])
		query, err := url.ParseQuery(matched[1])
		if err == nil {
			entry.ErrorType = "naxsi_exlog"
			entry.NaxsiMode = "fmt"
			entry.NaxsiMode = "exlog"

			entry.NaxsiExLogIP = query.Get("ip")
			entry.NaxsiExLogServer = query.Get("server")
			entry.NaxsiExLogURI = query.Get("uri")
			entry.NaxsiExLogID = query.Get("id")
			entry.NaxsiExLogZone = query.Get("zone")
			entry.NaxsiExLogVarName = query.Get("var_name")
			entry.NaxsiExLogContent = query.Get("content")

			entry.checkSumParts = []string{
				"naxsi_exlog",
				query.Get("server"),
				query.Get("uri"),
				query.Get("id"),
				query.Get("zone"),
				query.Get("var_name"),
				query.Get("content"),
			}

			entry.checkSumUseMsg = false
		}
	}
}
