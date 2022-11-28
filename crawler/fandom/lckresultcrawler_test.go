package fandom_test

import (
	"fmt"
	"testing"

	"github.com/gyu-young-park/lck_data_generator/crawler/fandom"
)

func TestLckResultCrawler(t *testing.T) {
	crawler := fandom.NewLCKSetResultCrawler()
	data := crawler.GetResult()
	fmt.Println(data)
}
