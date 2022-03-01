package main

import (
	"regexp"
)

var errorFastCGIRegex = regexp.MustCompile(`FastCGI sent in stderr: "(?P<fastCGI>[^"]+)"`)

func findErrorFastCGI(entry *nginxErrorEntry) {
	if ok := errorFastCGIRegex.MatchString(entry.Message); ok {
		matched := errorFastCGIRegex.FindStringSubmatch(entry.Message)
		entry.ErrorType = errorTypeFastcgi
		entry.ErrorDetails = stringPointer(matched[1])
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))

		entry.checkSumParts = []string{entry.ErrorType, *entry.ErrorDetails}

		entry.checkSumUseMsg = false
	}
}
