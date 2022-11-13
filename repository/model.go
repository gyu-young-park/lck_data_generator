package repository

import playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"

type LCKMatchModel struct {
	IsError     bool                              `json:"error"`
	PlayList    string                            `json:"playlist"`
	Title       string                            `json:"title"`
	VideoId     string                            `json:"video_id"`
	Season      string                            `json:"season"`
	Team1       string                            `json:"team1"`
	Outcome1    string                            `json:"outcome1"`
	Team2       string                            `json:"team2"`
	Outcome2    string                            `json:"outcome2"`
	WinTeam     string                            `json:"win_team"`
	LoseTeam    string                            `json:"lose_team"`
	Views       string                            `json:"views"`
	Thumbnails  playlistitems.VideoThumbnailModel `json:"thumbnails"`
	Date        string                            `json:"date"`
	PublishedAt int64                             `json:"published_at"`
}

func (l *LCKMatchModel) SetLCKMatchVideo(playlist string, title string, videoId string, season string, views string, thumbnails playlistitems.VideoThumbnailModel, date string, publishedAt int64) {
	l.PlayList = playlist
	l.Title = title
	l.VideoId = videoId
	l.Season = season
	l.Views = views
	l.Thumbnails = thumbnails
	l.Date = date
	l.PublishedAt = publishedAt
}

func (l *LCKMatchModel) SetLCKMatchScore(team1, score1, team2, score2 string) {
	l.Team1 = team1
	l.Outcome1 = score1
	l.Team2 = team2
	l.Outcome2 = score2
	if l.Outcome1 == "W" {
		l.WinTeam = team1
		l.LoseTeam = team2
	} else if l.Outcome1 == "L" {
		l.LoseTeam = team1
		l.WinTeam = team2
	}
}

type LCKMatchListModel struct {
	Data  []LCKMatchModel `json:"data"`
	Error string          `json:"error"`
}

type LCKTeamWithSeasonModel struct {
	Season   string   `json:"season"`
	TeamList []string `json:"teams"`
}

type LCKTeamWithSeasonListModel struct {
	Data  []LCKTeamWithSeasonModel `json:"data"`
	Error string                   `json:"error"`
}

type LCKSeasonWithTeamModel struct {
	Team       string   `json:"team"`
	SeasonList []string `json:"seasons"`
}

type LCKSeasonWithTeamListModel struct {
	Data  []LCKSeasonWithTeamModel `json:"data"`
	Error string                   `json:"error"`
}
