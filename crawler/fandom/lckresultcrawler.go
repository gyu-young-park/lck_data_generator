package fandom

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type LCKSetResultCrawler struct {
	collector   *colly.Collector
	queryOption *QueryOption
	result      []*LCKSetDataModel
	done        chan bool
	testResult  string
}

func NewLCKSetResultCrawler() *LCKSetResultCrawler {
	lckSetResultCrawler := &LCKSetResultCrawler{
		collector:   colly.NewCollector(),
		queryOption: NewQueryOption(),
	}
	lckSetResultCrawler.Ready()
	return lckSetResultCrawler
}

func (l *LCKSetResultCrawler) Ready() {
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

func (l *LCKSetResultCrawler) SetData(data interface{}) {

}

func (l *LCKSetResultCrawler) GetResult() interface{} {
	l.collector.UserAgent = "Mozilla/5.0"
	fmt.Printf("crawler-url: %s\n", l.queryOption.makeQueryUrl())
	err := l.collector.Visit(l.queryOption.makeQueryUrl())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return l.testResult
}
