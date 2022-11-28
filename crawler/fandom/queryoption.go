package fandom

import (
	"bytes"
	"fmt"
	"time"
)

// https://lol.fandom.com/wiki/Special:RunQuery/MatchHistoryGame?MHG%5Bpreload%5D=Tournament&MHG%5Btournament%5D=LCK+2019+Spring&MHG%5Bteam%5D=&MHG%5Bteam1%5D=&MHG%5Bteam2%5D=&MHG%5Bban%5D=&MHG%5Brecord%5D=&MHG%5Bascending%5D%5Bis_checkbox%5D=true&MHG%5Blimit%5D=&MHG%5Boffset%5D=&MHG%5Bregion%5D=&MHG%5Byear%5D=&MHG%5Bstartdate%5D=2019-03-21&MHG%5Benddate%5D=2019-03-22&MHG%5Bwhere%5D=&MHG%5Btextonly%5D%5Bis_checkbox%5D=true&_run=&pfRunQueryFormName=MatchHistoryGame&wpRunQuery=&pf_free_text=
const FANDOM_URL = "https://lol.fandom.com/wiki/Special:RunQuery/MatchHistoryGame?MHG%5Btextonly%5D=Yes&MHG%5Bpreload%5D=Tournament&MHG%5Bspl%5D=yes&_run=&"
const DEFAULT_SEASON = "LCK+2019+Spring"
const DEFAULT_DATE = "2019-03-21"
const DEFAULT_QUERY_PREFIX = "MHG%5B"
const DEFAULT_QUERY_SUFFIX = "%5D="
const DEFAULT_QUERY_TOURNAMENT = "tournament"
const DEFAULT_QUERY_STARTDATE = "startdate"
const DEFAULT_QUERY_ENDDATE = "enddate"

// https://lol.fandom.com/Special:RunQuery/MatchHistoryGame?MHG%5Benddate%5D=2019-03-23&MHG%5Btournament%5D=LCK%202019%20Spring&MHG%5Bstartdate%5D=2019-03-21&MHG%5Bpreload%5D=Tournament&MHG%5Bspl%5D=yes&_run=
// https://lol.fandom.com/wiki/Special:RunQuery/MatchHistoryGame?MHG%25%21B%28string=2019-03-22%29tournament%25%21D%28MISSING%29%3D%25%21s%28MISSING%29&MHG%25%21B%28MISSING%29team%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29team1%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29team2%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29ban%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29record%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29ascending%25%21D%28MISSING%29%25%21B%28MISSING%29is_checkbox%25%21D%28MISSING%29=true&MHG%25%21B%28MISSING%29limit%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29offset%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29region%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29year%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29startdate%25%21D%28MISSING%29=%25%21s%28MISSING%29&MHG%25%21B%28MISSING%29enddate%25%21D%28MISSING%29=%25%21s%28MISSING%29&MHG%25%21B%28MISSING%29where%25%21D%28MISSING%29=&MHG%25%21B%28MISSING%29textonly%25%21D%28MISSING%29%25%21B%28MISSING%29is_checkbox%25%21D%28MISSING%29=true&_run=&pfRunQueryFormName=MatchHistoryGame&wpRunQuery=&pf_free_text=
type QueryOption struct {
	url    string
	season string
	date   string
}

func NewQueryOption() *QueryOption {
	return &QueryOption{
		url:    FANDOM_URL,
		season: DEFAULT_SEASON,
		date:   DEFAULT_DATE,
	}
}

func (q *QueryOption) makeQueryParam(key, value string) string {
	var buf bytes.Buffer
	buf.WriteString(DEFAULT_QUERY_PREFIX)
	buf.WriteString(key)
	buf.WriteString(DEFAULT_QUERY_SUFFIX)
	buf.WriteString(value)
	return buf.String()
}

func (q *QueryOption) makeAllQueryString() string {
	var buf bytes.Buffer
	buf.WriteString(q.makeQueryParam(DEFAULT_QUERY_TOURNAMENT, q.season))
	buf.WriteString("&")
	buf.WriteString(q.makeQueryParam(DEFAULT_QUERY_STARTDATE, q.date))
	buf.WriteString("&")

	date, err := time.Parse("2006-01-02", q.date)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	nextDate := date.AddDate(0, 0, 1)
	nextDateString := nextDate.Format("2006-01-02")
	fmt.Println("next day:", nextDateString)

	buf.WriteString(q.makeQueryParam(DEFAULT_QUERY_ENDDATE, nextDateString))
	return buf.String()
}

func (q *QueryOption) makeQueryUrl() string {
	var buf bytes.Buffer
	buf.WriteString(q.url)
	buf.WriteString(q.makeAllQueryString())
	return buf.String()
}
