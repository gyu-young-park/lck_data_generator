package videoitem

import (
	"time"

	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

type VideoItemList []*VideoItem
type VideoItemListMapper map[string]VideoItemList

func NewVideoItemListMapper() VideoItemListMapper {
	ret := make(VideoItemListMapper)
	return ret
}

type VideoItem struct {
	PlayList    string
	Title       string
	VideoId     string
	Season      string
	Statistics  videostatistics.VideoStatisticsModel
	Thumbnails  playlistitems.VideoThumbnailModel
	PublishedAt time.Time
}

func NewVideoItem(playlist string, title string, videoId string, season string, statistics videostatistics.VideoStatisticsModel, thumbnails playlistitems.VideoThumbnailModel, publishedAt time.Time) *VideoItem {
	return &VideoItem{PlayList: playlist, Title: title, VideoId: videoId, Season: season, Statistics: statistics, Thumbnails: thumbnails, PublishedAt: publishedAt}
}
