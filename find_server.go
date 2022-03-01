package main

import (
	"regexp"
)

var serverRegex = regexp.MustCompile(`server[: ]+(?P<server>[^ ,]+)`)

func findServer(entry *nginxErrorEntry) {
	if ok := serverRegex.MatchString(entry.Message); ok {
		matched := serverRegex.FindStringSubmatch(entry.Message)
		entry.Server = stringPointer(matched[1])
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))
	}
}
