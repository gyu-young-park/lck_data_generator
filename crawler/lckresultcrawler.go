package crawler

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type LCKSetResultCrawler struct {
	collector   *colly.Collector
	QueryOption *InvenLCKResultQueryParam
	Result      []*LCKSetDataModel
	done        chan bool
}

func NewLCKSetResultCrawler() *LCKSetResultCrawler {
	lckResultCrawler := &LCKSetResultCrawler{
		collector: colly.NewCollector(),
		QueryOption: NewInvenLCKResultQueryParam(
			INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER,
			string(INVEN_DEFAULT_SHIP_GROUP),
			INVEN_LCK_DEFAULT_DATE),
	}
	lckResultCrawler.Ready()
	return lckResultCrawler
}

func (l *LCKSetResultCrawler) Ready() {
	l.collector.OnHTML("body > div#esportsBody div.listFrame", func(e *colly.HTMLElement) {
		leftTeamName := e.ChildText("div.wTeam > div.leftPart > a.teamname")
		leftTeamScore := e.ChildText("div.wTeam > div.rightPart > div")
		rightTeamScore := e.ChildText("div.lTeam > div.leftPart > div")
		rightTeamName := e.ChildText("div.lTeam > div.rightPart > a.teamname")
		l.Result = append(l.Result, &LCKSetDataModel{
			Date: l.QueryOption.Date,
			TeamScore1: TeamScore{
				Team:  leftTeamName,
				Score: leftTeamScore,
			},
			TeamScore2: TeamScore{
				Team:  rightTeamName,
				Score: rightTeamScore,
			},
		})
	})
}

func (l *LCKSetResultCrawler) Clear() {
	l.Result = l.Result[:0]
}

func (l *LCKSetResultCrawler) SetData(data interface{}) {
	l.Clear()
	date, _ := data.(string)
	l.QueryOption.Date = date
}

func (l *LCKSetResultCrawler) GetResult() interface{} {
	l.collector.Visit(l.QueryOption.MakeQueryURL())
	fmt.Printf("craowler-url: %s\n", l.QueryOption.MakeQueryURL())
	return l.Result
}

func (l *LCKSetResultCrawler) GoroutineGetResult(done <-chan interface{}) chan interface{} {
	resultDataChan := make(chan interface{})
	crawData := func(done <-chan interface{}) {
		l.collector.Visit(l.QueryOption.MakeQueryURL())
		fmt.Printf("craowler-url: %s\n", l.QueryOption.MakeQueryURL())
		select {
		case <-done:
			close(resultDataChan)
			return
		case resultDataChan <- l.Result:
		}
	}
	go crawData(done)
	return resultDataChan
}
