package fandom

type LCKSetDataModel struct {
	Data  string
	Team1 TeamInfo
	Team2 TeamInfo
}

type TeamInfo struct {
	Team      string
	Score     string
	Champions []string
	Roster    []string
}
