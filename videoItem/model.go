package videoitem

import (
	"time"

	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
)

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
	Thumbnails playlistitems.VideoThumbnailModel
	Date time.Time
}

func NewVideoItem(playlist string, title string, videoId string, season string, thumbnails playlistitems.VideoThumbnailModel,date time.Time) *VideoItem {
	return &VideoItem{PlayList: playlist, Title: title, VideoId: videoId, Season: season, Thumbnails: thumbnails,Date: date}
}