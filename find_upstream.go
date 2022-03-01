package main

import (
	"net/url"
	"regexp"
)

var upstreamRegex = regexp.MustCompile(`upstream[: ]+"(?P<upstream>[^"]+)"`)

func findUpstream(entry *nginxErrorEntry) {
	if ok := upstreamRegex.MatchString(entry.Message); ok {
		matched := upstreamRegex.FindStringSubmatch(entry.Message)
		entry.Upstream = stringPointer(matched[1])
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))

		u, err := url.Parse(matched[1])
		if err == nil {
			entry.UpstreamHost = stringPointer(u.Host)
		}
	}
}
