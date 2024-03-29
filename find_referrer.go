package main

import (
	"net/url"
	"regexp"
)

var referrerRegex = regexp.MustCompile(`referrer[: ]+"(?P<referrer>[^"]+)"`)

func findReferrer(entry *nginxErrorEntry) {
	if ok := referrerRegex.MatchString(entry.Message); ok {
		matched := referrerRegex.FindStringSubmatch(entry.Message)
		entry.Referrer = &matched[1]
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))

		u, err := url.Parse(*entry.Referrer)
		if err == nil {
			entry.ReferrerHost = &u.Host
		}
	}
}
