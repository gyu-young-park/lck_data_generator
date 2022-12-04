package videoitem

import (
	"errors"
	"fmt"
	"time"

	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

var (
	ErrDuplicatedData = errors.New("Error: Already data is exist")
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
	Statistics  videostatistics.VideoStatisticsResponseModel
	Thumbnails  playlistitems.VideoThumbnailModel
	PublishedAt time.Time
}

func NewVideoItem(playlist string, title string, videoId string, season string, statistics videostatistics.VideoStatisticsResponseModel, thumbnails playlistitems.VideoThumbnailModel, publishedAt time.Time) *VideoItem {
	return &VideoItem{PlayList: playlist, Title: title, VideoId: videoId, Season: season, Statistics: statistics, Thumbnails: thumbnails, PublishedAt: publishedAt}
}

func (v VideoItemListMapper) AppendWithDuplicatedCheck(key string, data *VideoItem) error {
	for _, videoItem := range v[key] {
		if videoItem.VideoId == data.VideoId {
			return fmt.Errorf("[%s] title:%s videoId:%s", ErrDuplicatedData.Error(), data.Title, data.VideoId)
		}
	}
	v[key] = append(v[key], data)
	return nil
}
