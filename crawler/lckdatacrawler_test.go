package crawler_test

import (
	"fmt"
	"testing"

	"github.com/gyu-young-park/lck_data_generator/crawler"
	"github.com/gyu-young-park/lck_data_generator/crawler/fandom"
)

func TestLckResultCrawler(t *testing.T) {
	crawlerImpl := fandom.NewLCKDataCrawler()
	data := crawlerImpl.GetResult().([]*crawler.LCKSetDataModel)
	for _, datum := range data {
		fmt.Println(datum)
	}
}
