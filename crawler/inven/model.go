package inven

type LCKSetDataModel struct {
	Date       string
	TeamScore1 TeamScore
	TeamScore2 TeamScore
}

type TeamScore struct {
	Team  string
	Score string
}
