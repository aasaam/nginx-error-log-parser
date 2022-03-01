package main

import (
	"net/url"
	"regexp"
	"strconv"
)

var naxsiFMTRegex = regexp.MustCompile(`NAXSI_FMT: (?P<naxsiFMT>[^ ]+),`)

var naxsiFmtCScoreRegex = regexp.MustCompile(`^cscore(?P<index>[0-9]+)$`)

func findNaxsiFmt(entry *nginxErrorEntry) {
	if ok := naxsiFMTRegex.MatchString(entry.Message); ok {
		matched := naxsiFMTRegex.FindStringSubmatch(entry.Message)
		entry.Msg = stringPointer(replaceMatched(*entry.Msg, matched[0]))

		query, err := url.ParseQuery(matched[1])

		if err == nil {
			entry.ErrorType = errorTypeNaxsiFmt
			entry.NaxsiFmtIP = stringPointer(query.Get("ip"))
			entry.NaxsiFmtServer = stringPointer(query.Get("server"))
			entry.NaxsiFmtLearning = boolPointer(query.Get("learning") == "1")
			entry.NaxsiFmtBlock = boolPointer(query.Get("block") == "1")
			entry.NaxsiFmtVers = stringPointer(query.Get("vers"))
			entry.NaxsiFmtURI = stringPointer(query.Get("uri"))
			entry.NaxsiFmtTotalProcessed = intPointer(prettySureInt(query.Get("total_processed")))
			entry.NaxsiFmtTotalBlocked = intPointer(prettySureInt(query.Get("total_blocked")))
			entry.checkSumParts = []string{
				errorTypeNaxsiFmt,
				query.Get("server"),
				query.Get("uri"),
				query.Get("block"),
			}

			numberOfFmtItems := 0
			for key := range query {
				if ok := naxsiFmtCScoreRegex.MatchString(key); ok {
					numberOfFmtItems++
				}
			}

			for i := 0; i < numberOfFmtItems; i++ {
				item := naxsiFmtItem{}
				CScore := "cscore" + strconv.Itoa(i)
				_, CScoreFound := query[CScore]
				if CScoreFound {
					item.CScore = query.Get(CScore)
					entry.checkSumParts = append(entry.checkSumParts, item.CScore)
				}

				Score := "score" + strconv.Itoa(i)
				_, ScoreFound := query[Score]
				if ScoreFound {
					item.Score = query.Get(Score)
					entry.checkSumParts = append(entry.checkSumParts, item.Score)
				}

				ID := "id" + strconv.Itoa(i)
				_, IDFound := query[ID]
				if IDFound {
					item.ID = query.Get(ID)
					entry.checkSumParts = append(entry.checkSumParts, item.ID)
				}

				VarName := "var_name" + strconv.Itoa(i)
				_, VarNameFound := query[VarName]
				if VarNameFound {
					item.VarName = query.Get(VarName)
					entry.checkSumParts = append(entry.checkSumParts, item.VarName)
				}

				Zone := "zone" + strconv.Itoa(i)
				_, ZoneFound := query[Zone]
				if ZoneFound {
					item.Zone = query.Get(Zone)
					entry.checkSumParts = append(entry.checkSumParts, item.Zone)
				}
				entry.NaxsiFmtItems = append(entry.NaxsiFmtItems, item)
			}

			entry.checkSumUseMsg = false
		}
	}
}
