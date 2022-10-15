package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

type LCKResultCrawler struct {
	collector   *colly.Collector
	QueryOption *InvenLCKResultQueryParam
}

func NewLCKResultCrawler() *LCKResultCrawler {
	lckResultCrawler := &LCKResultCrawler{
		collector: colly.NewCollector(),
		QueryOption: NewInvenLCKResultQueryParam(
			INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER,
			string(INVEN_DEFAULT_SHIP_GROUP),
			INVEN_LCK_DEFAULT_DATE),
	}
	lckResultCrawler.Ready()
	return lckResultCrawler
}

func (l *LCKResultCrawler) Ready() {
	l.collector.OnHTML("body > div#esportsBody div.wTeam a.teamname", func(e *colly.HTMLElement) {
		fmt.Println("---------------cralwer left start--------------")
		fmt.Println(e.Text)
	})

	l.collector.OnHTML("body > div#esportsBody div.wTeam div.rightPart", func(e *colly.HTMLElement) {
		fmt.Println("---------------cralwer left start--------------")
		fmt.Println(e.Text)
	})

	l.collector.OnHTML("body > div#esportsBody div.lTeam a.teamname", func(e *colly.HTMLElement) {
		fmt.Println("---------------cralwer right start--------------")
		fmt.Println(e.Text)
	})

	l.collector.OnHTML("body > div#esportsBody div.lTeam div.leftPart", func(e *colly.HTMLElement) {
		fmt.Println("---------------cralwer right start--------------")
		fmt.Println(strings.TrimSpace(e.Text))
	})
}

func (l *LCKResultCrawler) SetData(data interface{}) {
	date, _ := data.(string)
	l.QueryOption.Date = date
}

func (l *LCKResultCrawler) GetResult() {
	l.collector.Visit(l.QueryOption.MakeQueryURL())
}
