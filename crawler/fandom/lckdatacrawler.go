package fandom

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	crawlermodel "github.com/gyu-young-park/lck_data_generator/crawler/model"
)

type LCKDataCrawler struct {
	collector   *colly.Collector
	queryOption *crawlermodel.QueryOption
	result      []*crawlermodel.LCKSetDataModel
}

func NewLCKDataCrawler() *LCKDataCrawler {
	lckSetResultCrawler := &LCKDataCrawler{
		collector:   colly.NewCollector(),
		queryOption: NewQueryOption(),
	}
	lckSetResultCrawler.Ready()
	return lckSetResultCrawler
}

func (l *LCKDataCrawler) Ready() {
	l.collector.OnHTML(" div.wide-content-scroll", func(e *colly.HTMLElement) {
		e.ForEach("tbody > tr.multirow-highlighter", func(i int, element *colly.HTMLElement) {
			var resultData crawlermodel.LCKSetDataModel
			patch := element.ChildText("td:nth-child(2)")
			fmt.Println(i, " ", patch)
			resultData.Team1.Team = element.ChildText("td:nth-child(3)")
			resultData.Team2.Team = element.ChildText("td:nth-child(4)")
			winner := element.ChildText("td:nth-child(5)")
			if winner == resultData.Team1.Team {
				resultData.Team1.Score = "W"
				resultData.Team2.Score = "L"
			} else {
				resultData.Team1.Score = "L"
				resultData.Team2.Score = "W"
			}
			resultData.Team1.Champions = element.ChildText("td:nth-child(8)")
			resultData.Team2.Champions = element.ChildText("td:nth-child(9)")
			resultData.Team1.Roster = element.ChildText("td:nth-child(10)")
			resultData.Team2.Roster = element.ChildText("td:nth-child(11)")
			resultData.Date = l.queryOption.Date
			l.result = append(l.result, &resultData)
		})
	})

	l.collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
}

func (l *LCKDataCrawler) Clear() {
	l.result = l.result[:0]
	l.queryOption = nil
}

func (l *LCKDataCrawler) SetQueryOption(queryOption *crawlermodel.QueryOption) {
	l.Clear()
	l.queryOption = queryOption
}

func (l *LCKDataCrawler) GetResult() interface{} {
	l.collector.UserAgent = "Mozilla/5.0"
	fmt.Printf("crawler-url: %s\n", makeQueryUrl(l.queryOption))
	err := l.collector.Visit(makeQueryUrl(l.queryOption))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return l.result
}
