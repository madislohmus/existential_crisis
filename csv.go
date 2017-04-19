package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
)

type card struct {
	Description string
	Self        points
	Friend      points
	All         points
	Crisis      bool
}

type points struct {
	Time   int64
	Social int64
	Health int64
	Money  int64
}

func readFile(filename string) ([]card, []card, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	reader := csv.NewReader(bufio.NewReader(file))

	var crisisCards []card
	var normalCards []card
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	for _, record := range records[3:] {
		if len(record[0]) == 0 {
			continue
		}
		c := card{}
		c.Description = record[0]

		if len(record[1]) > 0 {
			c.Crisis = true
		}

		sTimeStr := record[2]
		sSocialStr := record[3]
		sHealthStr := record[4]
		sMoneyStr := record[5]

		selfPoints, err := getPoints(sTimeStr, sSocialStr, sHealthStr, sMoneyStr)
		if err != nil {
			return nil, nil, err
		}
		c.Self = *selfPoints

		fSocialStr := record[6]
		fHealthStr := record[7]
		fMoneyStr := record[8]
		fTimeStr := record[9]

		friendPoints, err := getPoints(fTimeStr, fSocialStr, fHealthStr, fMoneyStr)
		if err != nil {
			return nil, nil, err
		}
		c.Friend = *friendPoints

		aSocialStr := record[10]
		aHealthStr := record[11]
		aMoneyStr := record[12]
		aTimeStr := record[13]

		allPoints, err := getPoints(aTimeStr, aSocialStr, aHealthStr, aMoneyStr)
		if err != nil {
			return nil, nil, err
		}
		c.All = *allPoints
		if c.Crisis {
			crisisCards = append(crisisCards, c)
		} else {
			normalCards = append(normalCards, c)
		}
	}
	return normalCards, crisisCards, nil
}

func getPoints(sTimeStr, sSocialStr, sHealthStr, sMoneyStr string) (*points, error) {
	p := points{}
	if len(sTimeStr) == 0 {
		p.Time = 0
	} else {
		sTimeInt, err := strconv.ParseInt(sTimeStr, 10, 64)
		if err != nil {
			return nil, err
		}
		p.Time = sTimeInt
	}
	if len(sSocialStr) == 0 {
		p.Social = 0
	} else {
		sSocialInt, err := strconv.ParseInt(sSocialStr, 10, 64)
		if err != nil {
			return nil, err
		}
		p.Social = sSocialInt
	}
	if len(sHealthStr) == 0 {
		p.Health = 0
	} else {
		sHealthInt, err := strconv.ParseInt(sHealthStr, 10, 64)
		if err != nil {
			return nil, err
		}
		p.Health = sHealthInt
	}
	if len(sMoneyStr) == 0 {
		p.Money = 0
	} else {
		sMoneyInt, err := strconv.ParseInt(sMoneyStr, 10, 64)
		if err != nil {
			return nil, err
		}
		p.Money = sMoneyInt
	}
	return &p, nil
}
