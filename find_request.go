package main

import (
	"regexp"
)

var requestRegex = regexp.MustCompile(`request[: ]+"(?P<requestMethod>[A-Z]+) (?P<requestURI>.*) HTTP\/(?P<requestHTTPVersion>[0-9\.]+)"`)

func findRequest(entry *nginxErrorEntry) {
	if ok := requestRegex.MatchString(entry.Message); ok {
		matched := requestRegex.FindStringSubmatch(entry.Message)
		entry.RequestMethod = stringPointer(matched[1])
		entry.RequestURI = stringPointer(matched[2])
		entry.RequestHTTPVersion = stringPointer(matched[3])
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))
	}
}
