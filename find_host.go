package main

import (
	"regexp"
)

var hostRegex = regexp.MustCompile(`host[: ]+"(?P<host>[^"]+)"`)

func findHost(entry *nginxErrorEntry) {
	if ok := hostRegex.MatchString(entry.Message); ok {
		matched := hostRegex.FindStringSubmatch(entry.Message)
		entry.Host = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])
	}
}
