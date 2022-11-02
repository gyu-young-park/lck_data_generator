package videoitem

import "time"

type VideoItemList []*VideoItem
type VideoItemListMapper map[string]VideoItemList

func NewVideoItemListMapper() VideoItemListMapper{
	ret := make(VideoItemListMapper)
	return ret
}

type VideoItem struct{
	PlayList string
	Title string
	VideoId string
	Season string
	Date time.Time
}

func NewVideoItem(playlist string, title string, videoId string, season string,date time.Time) *VideoItem {
	return &VideoItem{PlayList: playlist, Title: title, VideoId: videoId, Season: season,Date: date}
}