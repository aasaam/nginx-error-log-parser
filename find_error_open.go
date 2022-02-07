package main

import (
	"regexp"
)

var errorOpenFailedRegex = regexp.MustCompile(`open\(\) "(?P<open>[^"]+)" failed \(2: No such file or directory\)`)

func findErrorOpenFailed(entry *nginxErrorEntry) {
	if ok := errorOpenFailedRegex.MatchString(entry.Message); ok {
		matched := errorOpenFailedRegex.FindStringSubmatch(entry.Message)
		entry.ErrorType = "open_failed"
		entry.ErrorDetails = matched[1]
		entry.Msg = replaceMatched(entry.Msg, matched[0])

		entry.checkSumParts = []string{entry.ErrorType, entry.ErrorDetails}

		entry.checkSumUseMsg = false
	}
}
