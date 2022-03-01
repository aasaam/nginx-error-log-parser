package main

import (
	"net/url"
	"regexp"
)

var naxsiExLogRegex = regexp.MustCompile(`NAXSI_EXLOG: (?P<naxsiExLog>[^ ]+),`)

func findNaxsiExLog(entry *nginxErrorEntry) {
	if ok := naxsiExLogRegex.MatchString(entry.Message); ok {
		matched := naxsiExLogRegex.FindStringSubmatch(entry.Message)
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))
		query, err := url.ParseQuery(matched[1])
		if err == nil {
			entry.ErrorType = errorTypeNaxsiExLog

			entry.NaxsiExLogIP = stringPointer(query.Get("ip"))
			entry.NaxsiExLogServer = stringPointer(query.Get("server"))
			entry.NaxsiExLogURI = stringPointer(query.Get("uri"))
			entry.NaxsiExLogID = stringPointer(query.Get("id"))
			entry.NaxsiExLogZone = stringPointer(query.Get("zone"))
			entry.NaxsiExLogVarName = stringPointer(query.Get("var_name"))
			entry.NaxsiExLogContent = stringPointer(query.Get("content"))

			entry.checkSumParts = []string{
				errorTypeNaxsiExLog,
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
