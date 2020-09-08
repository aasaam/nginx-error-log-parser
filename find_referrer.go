package main

import (
	"net/url"
	"regexp"
)

var referrerRegex, _ = regexp.Compile(`referrer[: ]+"(?P<referrer>[^"]+)"`)

func findReferrer(entry *NginxErrorEntry) {
	if ok := referrerRegex.MatchString(entry.Message); ok {
		matched := referrerRegex.FindStringSubmatch(entry.Message)
		entry.Referrer = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])

		u, err := url.Parse(entry.Referrer)
		if err == nil {
			entry.ReferrerHost = u.Host
		}
	}
}
