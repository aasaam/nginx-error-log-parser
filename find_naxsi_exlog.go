package main

import (
	"regexp"

	"github.com/derekstavis/go-qs"
)

var naxsiExLogRegex, _ = regexp.Compile(`NAXSI_EXLOG: (?P<naxsiExLog>[^ ]+),`)

func findNaxsiExLog(entry *NginxErrorEntry) {
	if ok := naxsiExLogRegex.MatchString(entry.Message); ok {
		matched := naxsiExLogRegex.FindStringSubmatch(entry.Message)
		entry.Msg = replaceMatched(entry.Msg, matched[0])
		query, e := qs.Unmarshal(matched[1])
		if e == nil {
			entry.NaxsiMode = "fmt"

			entry.NaxsiMode = "exlog"

			entry.NaxsiExLogIP = query["ip"].(string)
			entry.NaxsiExLogServer = query["server"].(string)
			entry.NaxsiExLogURI = query["uri"].(string)
			entry.NaxsiExLogID = query["id"].(string)
			entry.NaxsiExLogZone = query["zone"].(string)
			entry.NaxsiExLogVarName = query["var_name"].(string)
			entry.NaxsiExLogContent = query["content"].(string)

			entry.checkSumParts = []string{
				"naxsi_exlog",
				query["server"].(string),
				query["uri"].(string),
				query["id"].(string),
				query["zone"].(string),
				query["var_name"].(string),
				query["content"].(string),
			}

			entry.checkSumUseMsg = false
		}
	}
}
