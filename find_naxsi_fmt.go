package main

import (
	"regexp"
	"strconv"

	"github.com/derekstavis/go-qs"
)

var naxsiFMTRegex, _ = regexp.Compile(`NAXSI_FMT: (?P<naxsiFMT>[^ ]+),`)

var naxsiFmtCScoreRegex, _ = regexp.Compile(`^cscore(?P<index>[0-9]+)$`)

func findNaxsiFmt(entry *NginxErrorEntry) {
	if ok := naxsiFMTRegex.MatchString(entry.Message); ok {
		matched := naxsiFMTRegex.FindStringSubmatch(entry.Message)
		entry.Msg = replaceMatched(entry.Msg, matched[0])
		query, e := qs.Unmarshal(matched[1])
		if e == nil {
			entry.NaxsiMode = "fmt"
			entry.NaxsiFmtIP = query["ip"].(string)
			entry.NaxsiFmtServer = query["server"].(string)
			entry.NaxsiFmtLearning = query["learning"].(string) == "1"
			entry.NaxsiFmtBlock = query["block"].(string) == "1"
			entry.NaxsiFmtVers = query["vers"].(string)
			entry.NaxsiFmtURI = query["uri"].(string)
			entry.NaxsiFmtTotalProcessed, _ = strconv.Atoi(query["total_processed"].(string))
			entry.NaxsiFmtTotalBlocked, _ = strconv.Atoi(query["total_blocked"].(string))
			entry.checkSumParts = []string{
				"naxsi_fmt",
				query["server"].(string),
				query["uri"].(string),
				query["block"].(string),
			}

			numberOfFmtItems := 0
			for key := range query {
				if ok := naxsiFmtCScoreRegex.MatchString(key); ok {
					numberOfFmtItems++
				}

			}

			for i := 0; i < numberOfFmtItems; i++ {
				item := NaxsiFmtItem{}
				CScore := "cscore" + strconv.Itoa(i)
				_, CScoreFound := query[CScore]
				if CScoreFound {
					item.CScore = query[CScore].(string)
					entry.checkSumParts = append(entry.checkSumParts, item.CScore)
				}

				Score := "score" + strconv.Itoa(i)
				_, ScoreFound := query[Score]
				if ScoreFound {
					item.Score = query[Score].(string)
					entry.checkSumParts = append(entry.checkSumParts, item.Score)
				}

				ID := "id" + strconv.Itoa(i)
				_, IDFound := query[ID]
				if IDFound {
					item.ID = query[ID].(string)
					entry.checkSumParts = append(entry.checkSumParts, item.ID)
				}

				VarName := "var_name" + strconv.Itoa(i)
				_, VarNameFound := query[VarName]
				if VarNameFound {
					item.VarName = query[VarName].(string)
					entry.checkSumParts = append(entry.checkSumParts, item.VarName)
				}

				Zone := "zone" + strconv.Itoa(i)
				_, ZoneFound := query[Zone]
				if ZoneFound {
					item.Zone = query[Zone].(string)
					entry.checkSumParts = append(entry.checkSumParts, item.Zone)
				}
				entry.NaxsiFmtItems = append(entry.NaxsiFmtItems, item)
			}

			entry.checkSumUseMsg = false
		}
	}
}
