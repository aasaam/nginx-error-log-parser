package main

import (
	"regexp"
)

var serverRegex, _ = regexp.Compile(`server[: ]+(?P<server>[^ ,]+)`)

func findServer(entry *NginxErrorEntry) {
	if ok := serverRegex.MatchString(entry.Message); ok {
		matched := serverRegex.FindStringSubmatch(entry.Message)
		entry.Server = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])
	}
}
