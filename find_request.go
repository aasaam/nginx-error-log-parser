package main

import (
	"regexp"
)

var requestRegex = regexp.MustCompile(`request[: ]+"(?P<requestMethod>[A-Z]+) (?P<requestURI>.*) HTTP\/(?P<requestHTTPVersion>[0-9\.]+)"`)

func findRequest(entry *NginxErrorEntry) {
	if ok := requestRegex.MatchString(entry.Message); ok {
		matched := requestRegex.FindStringSubmatch(entry.Message)
		entry.RequestMethod = matched[1]
		entry.RequestURI = matched[2]
		entry.RequestHTTPVersion = matched[3]
		entry.Msg = replaceMatched(entry.Msg, matched[0])
	}
}
