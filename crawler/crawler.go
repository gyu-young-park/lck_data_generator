package crawler

import (
	"github.com/gyu-young-park/lck_data_generator/crawler/fandom"
	"github.com/gyu-young-park/lck_data_generator/crawler/inven"
	crawlermodel "github.com/gyu-young-park/lck_data_generator/crawler/model"
)

type CrawlerMode_t string

var (
	CRAWLER_FANDOM_MODE = CrawlerMode_t("fandom")
	CRAWLER_INVEN_MODE = CrawlerMode_t("inven")
)

type Crawler interface {
	Ready()
	SetQueryOption(queryOption *crawlermodel.QueryOption)
	GetResult() interface{}
}

func New(mode CrawlerMode_t) Crawler {
	switch mode {
	case CRAWLER_INVEN_MODE:
		return inven.NewLCKSetResultCrawler()
	case CRAWLER_FANDOM_MODE:
		return fandom.NewLCKDataCrawler()
	default:
		return nil
	}
}