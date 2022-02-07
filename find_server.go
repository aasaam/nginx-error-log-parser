package main

import (
	"regexp"
)

var serverRegex = regexp.MustCompile(`server[: ]+(?P<server>[^ ,]+)`)

func findServer(entry *nginxErrorEntry) {
	if ok := serverRegex.MatchString(entry.Message); ok {
		matched := serverRegex.FindStringSubmatch(entry.Message)
		entry.Server = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])
	}
}
