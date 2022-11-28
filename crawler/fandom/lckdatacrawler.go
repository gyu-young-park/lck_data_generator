package fandom

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type LCKDataCrawler struct {
	collector   *colly.Collector
	queryOption *QueryOption
	result      []*LCKSetDataModel
	testResult  string
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
			patch := element.ChildText("td:nth-child(2)")
			fmt.Println(i, " ", patch)
			team1 := element.ChildText("td:nth-child(3)")
			fmt.Println(i, " ", team1)
		})
	})

	l.collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})
}

func (l *LCKDataCrawler) SetData(data interface{}) {

}

func (l *LCKDataCrawler) GetResult() interface{} {
	l.collector.UserAgent = "Mozilla/5.0"
	fmt.Printf("crawler-url: %s\n", l.queryOption.makeQueryUrl())
	err := l.collector.Visit(l.queryOption.makeQueryUrl())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return l.testResult
}
