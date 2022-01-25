package main

import (
	"net/url"
	"regexp"
)

var upstreamRegex = regexp.MustCompile(`upstream[: ]+"(?P<upstream>[^"]+)"`)

func findUpstream(entry *NginxErrorEntry) {
	if ok := upstreamRegex.MatchString(entry.Message); ok {
		matched := upstreamRegex.FindStringSubmatch(entry.Message)
		entry.Upstream = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])

		u, err := url.Parse(entry.Upstream)
		if err == nil {
			entry.UpstreamHost = u.Host
		}
	}
}
