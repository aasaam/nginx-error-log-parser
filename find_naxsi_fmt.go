package main

import (
	"net/url"
	"regexp"
	"strconv"
)

var naxsiFMTRegex = regexp.MustCompile(`NAXSI_FMT: (?P<naxsiFMT>[^ ]+),`)

var naxsiFmtCScoreRegex = regexp.MustCompile(`^cscore(?P<index>[0-9]+)$`)

func findNaxsiFmt(entry *NginxErrorEntry) {
	if ok := naxsiFMTRegex.MatchString(entry.Message); ok {
		matched := naxsiFMTRegex.FindStringSubmatch(entry.Message)
		entry.Msg = replaceMatched(entry.Msg, matched[0])

		query, err := url.ParseQuery(matched[1])

		if err == nil {
			entry.ErrorType = "naxsi_fmt"
			entry.NaxsiMode = "fmt"
			entry.NaxsiFmtIP = query.Get("ip")
			entry.NaxsiFmtServer = query.Get("server")
			entry.NaxsiFmtLearning = query.Get("learning") == "1"
			entry.NaxsiFmtBlock = query.Get("block") == "1"
			entry.NaxsiFmtVers = query.Get("vers")
			entry.NaxsiFmtURI = query.Get("uri")
			entry.NaxsiFmtTotalProcessed, _ = strconv.Atoi(query.Get("total_processed"))
			entry.NaxsiFmtTotalBlocked, _ = strconv.Atoi(query.Get("total_blocked"))
			entry.checkSumParts = []string{
				"naxsi_fmt",
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
				item := NaxsiFmtItem{}
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
