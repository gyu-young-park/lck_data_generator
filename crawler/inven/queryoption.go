package inven

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/crawler"
)

type SHIPGROUP string

const (
	ALL   = SHIPGROUP("")
	LCK   = SHIPGROUP("1")
	CHAMS = SHIPGROUP("3")
)

const INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER = "https://lol.inven.co.kr/dataninfo/match/teamList.php?iskin=esports&category=&category2=&shipcode=&shipgroup=%s&teamName=&teamName2=&startDate=%s&endDate=%s"
const INVEN_DEFAULT_SHIP_GROUP = LCK
const INVEN_LCK_DEFAULT_DATE = "2019-01-18"

func newInvenLCKResultQueryParam(url string, shipGroup string, date string) *crawler.QueryOption {
	return &crawler.QueryOption{
		Url:              url,
		InvenQueryOption: crawler.InvenQueryOption{ShipGroup: shipGroup},
		Date:             date,
	}
}

func NewInvenLCKResultQueryParamWithDate(date string) *crawler.QueryOption {
	return &crawler.QueryOption{
		Url:              INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER,
		InvenQueryOption: crawler.InvenQueryOption{ShipGroup: string(ALL)},
		Date:             date,
	}
}

func makeQueryURL(queryOption *crawler.QueryOption) string {
	return fmt.Sprintf(queryOption.Url, ALL, queryOption.Date, queryOption.Date)
}
