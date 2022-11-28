package inven_test

import (
	"fmt"
	"testing"

	"github.com/gyu-young-park/lck_data_generator/crawler/inven"
)

func TestLckResultCrawler(t *testing.T) {
	crawler := inven.NewLCKSetResultCrawler()
	data := crawler.GetResult()
	fmt.Println(data)
}
