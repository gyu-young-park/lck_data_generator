package crawler

import "fmt"

type SHIPGROUP string

const (
	ALL   = SHIPGROUP("")
	LCK   = SHIPGROUP("1")
	CHAMS = SHIPGROUP("3")
)

const INVEN_LCK_CAHMPIONSHIP_RESULT_URL_FORMATTER = "https://lol.inven.co.kr/dataninfo/match/teamList.php?iskin=esports&category=&category2=&shipcode=&shipgroup=%s&teamName=&teamName2=&startDate=%s&endDate=%s"
const INVEN_DEFAULT_SHIP_GROUP = LCK
const INVEN_LCK_DEFAULT_DATE = "2019-01-18"

type InvenLCKResultQueryParam struct {
	Url       string
	ShipGroup string
	Date      string
}

func NewInvenLCKResultQueryParam(url string, shipGroup string, date string) *InvenLCKResultQueryParam {
	return &InvenLCKResultQueryParam{
		Url:       url,
		ShipGroup: shipGroup,
		Date:      date,
	}
}

func (i *InvenLCKResultQueryParam) SetDate(date string) {
	i.Date = date
}

func (i *InvenLCKResultQueryParam) SetShipGroup(shipGroup string) {
	i.ShipGroup = shipGroup
}

func (i *InvenLCKResultQueryParam) MakeQueryURL() string {
	return fmt.Sprintf(i.Url, ALL, i.Date, i.Date)
}

func (i *InvenLCKResultQueryParam) MakeQueryURLWithDate(date string) string {
	return fmt.Sprintf(i.Url, ALL, date, date)
}
