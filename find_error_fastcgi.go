package main

import (
	"regexp"
)

var errorFastCGIRegex, _ = regexp.Compile(`FastCGI sent in stderr: "(?P<fastCGI>[^"]+)"`)

func findErrorFastCGI(entry *NginxErrorEntry) {
	if ok := errorFastCGIRegex.MatchString(entry.Message); ok {
		matched := errorFastCGIRegex.FindStringSubmatch(entry.Message)
		entry.ErrorType = "fastcgi_error"
		entry.ErrorDetails = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])

		entry.checkSumParts = []string{entry.ErrorType, entry.ErrorDetails}

		entry.checkSumUseMsg = false
	}
}
