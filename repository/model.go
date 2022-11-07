package repository

import playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"

type LCKMatchTeamModel struct {
	Team1    string `json:"team1"`
	Outcome1 string `json:"outcome1"`
	Team2    string `json:"team2"`
	Outcome2 string `json:"outcome2"`
}

func NewLCKMatchTeamModel(team1, outcome1, team2, outcome2 string) *LCKMatchTeamModel {
	return &LCKMatchTeamModel{
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
	Season   string `json:"season"`
	Thumbnails playlistitems.VideoThumbnailModel `json:"thumbnails"`
	Date     string `json:"date"`
}

func NewLCKMatchVideoModel(playlist, title, videoId ,season string, thumbnails playlistitems.VideoThumbnailModel,date string) *LCKMatchVideoModel {
	return &LCKMatchVideoModel{
		PlayList: playlist,
		Title:    title,
		VideoId:  videoId,
		Season: season,
		Thumbnails: thumbnails,
		Date:     date,
	}
}

type LCKMatchModel struct {
	IsError bool `json:"error"`
	LCKMatchVideoModel
	LCKMatchTeamModel
}

type LCKMatchListModel struct {
	Data  []LCKMatchModel `json:"data"`
	Error string          `json:"error"`
}

type LCKTeamWithSeasonModel struct {
	Season string `json:"season"`
	TeamList []string `json:"teams"`
}

type LCKTeamWithSeasonListModel struct {
	Data  []LCKTeamWithSeasonModel `json:"data"`
	Error string          `json:"error"`
}

type LCKSeasonWithTeamModel struct {
	Team string `json:"team"`
	SeasonList []string `json:"seasons"`
}

type LCKSeasonWithTeamListModel struct {
	Data  []LCKSeasonWithTeamModel `json:"data"`
	Error string          `json:"error"`
}