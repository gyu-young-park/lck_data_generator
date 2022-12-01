package inven

import (
	"fmt"

	crawlermodel "github.com/gyu-young-park/lck_data_generator/crawler/model"
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

func newInvenLCKResultQueryParam(url string, shipGroup string, date string) *crawlermodel.QueryOption {
	return &crawlermodel.QueryOption{
		Url:              url,
		InvenQueryOption: crawlermodel.InvenQueryOption{ShipGroup: shipGroup},
		Date:             date,
	}
}

func NewInvenLCKResultQueryParamWithDate(date string) *crawlermodel.QueryOption {
	return &crawlermodel.QueryOption{
		Url:              INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER,
		InvenQueryOption: crawlermodel.InvenQueryOption{ShipGroup: string(ALL)},
		Date:             date,
	}
}

func makeQueryURL(queryOption *crawlermodel.QueryOption) string {
	return fmt.Sprintf(queryOption.Url, ALL, queryOption.Date, queryOption.Date)
}
