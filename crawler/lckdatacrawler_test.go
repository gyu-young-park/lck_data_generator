package crawler_test

import (
	"fmt"
	"testing"
	"time"

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

func TestTimeCompare(t *testing.T) {
	rawtime1 := "2019-03-23"
	time1, err := time.Parse("2006-01-02", rawtime1)
	nextDate := time1.AddDate(0, 0, 1)
	preDate := time1.AddDate(0, 0, -1)
	if err != nil {
		fmt.Println(err)
		return
	}
	before := time1.Before(nextDate)
	after := time1.After(preDate)
	fmt.Println(before, after)
}
