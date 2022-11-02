package repository

type LCKMathTeamModel struct {
	Team1    string `json:"team1"`
	Outcome1 string `json:"outcome1"`
	Team2    string `json:"team2"`
	Outcome2 string `json:"outcome2"`
}

func NewLCKMathTeamModel(team1, outcome1, team2, outcome2 string) *LCKMathTeamModel {
	return &LCKMathTeamModel{
		Team1:    team1,
		Team2:    team2,
		Outcome1: outcome1,
		Outcome2: outcome2,
	}
}

type LCKMatchVideoModel struct {
	PlayList string `json:"playlist"`
	Title    string `json:"title"`
	VideoId  string `json:"video_id"`
	Date     string `json:"date"`
}

func NewLCKMatchVideoModel(playlist, title, videoId, date string) *LCKMatchVideoModel {
	return &LCKMatchVideoModel{
		PlayList: playlist,
		Title:    title,
		VideoId:  videoId,
		Date:     date,
	}
}

type LCKMatchModel struct {
	IsError bool `json:"error"`
	LCKMatchVideoModel
	LCKMathTeamModel
}

type LCKMatchListModel struct {
	Data  []LCKMatchModel `json:"data"`
	Error string          `json:"error"`
}
