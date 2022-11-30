package inven

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gyu-young-park/lck_data_generator/crawler"
)

type LCKSetResultCrawler struct {
	collector   *colly.Collector
	queryOption *crawler.QueryOption
	result      []*crawler.LCKSetDataModel
	done        chan bool
}

func NewLCKSetResultCrawler() *LCKSetResultCrawler {
	lckResultCrawler := &LCKSetResultCrawler{
		collector: colly.NewCollector(),
		queryOption: newInvenLCKResultQueryParam(
			INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER,
			string(INVEN_DEFAULT_SHIP_GROUP),
			INVEN_LCK_DEFAULT_DATE),
	}
	lckResultCrawler.Ready()
	return lckResultCrawler
}

func (l *LCKSetResultCrawler) Ready() {
	l.collector.OnHTML("body > div#esportsBody div.listFrame", func(e *colly.HTMLElement) {
		var resultData crawler.LCKSetDataModel
		leftTeamName := e.ChildText("div.wTeam > div.leftPart > a.teamname")
		leftTeamScore := e.ChildText("div.wTeam > div.rightPart > div")
		rightTeamScore := e.ChildText("div.lTeam > div.leftPart > div")
		rightTeamName := e.ChildText("div.lTeam > div.rightPart > a.teamname")
		resultData.Date = l.queryOption.Date
		resultData.Team1.Team = leftTeamName
		resultData.Team1.Score = leftTeamScore
		resultData.Team2.Team = rightTeamName
		resultData.Team2.Score = rightTeamScore
		l.result = append(l.result, &resultData)
	})
}

func (l *LCKSetResultCrawler) Clear() {
	l.result = l.result[:0]
	l.queryOption = nil
}

func (l *LCKSetResultCrawler) SetQueryOption(queryOption *crawler.QueryOption) {
	l.Clear()
	l.queryOption = queryOption
}

func (l *LCKSetResultCrawler) GetResult() interface{} {
	l.collector.Visit(makeQueryURL(l.queryOption))
	fmt.Printf("crawler-url: %s\n", makeQueryURL(l.queryOption))
	return l.result
}
