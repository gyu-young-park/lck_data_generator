package calibratorvideo

import (
	"fmt"

	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

type VideoInjectorImpl struct {
}

func NewVideoInjectorImpl(crawlerMode string) *VideoInjectorImpl {
	setOmittedVideoCalibrateTable(crawlerMode)
	return &VideoInjectorImpl{}
}

func (v *VideoInjectorImpl) Calibrate(videoMapper videoitem.VideoItemListMapper, videoStatisticsService videostatistics.Service) {
	for date, videoList := range videoMapper {
		caliData, isExist := omittedVideoCalibrateTable[date]
		if !isExist {
			continue
		}
		// append videoList
		for _, videoId := range caliData.VideoIdList {
			// get video id
			videoStatisticsData, err := videoStatisticsService.GetVideoStatistics(videoId)
			fmt.Printf("Omitted Video Calibrated: %s %s %s\n", videoId, caliData.PlayListTitle, caliData.Date)
			videoItem := videoitem.NewVideoItem(
				caliData.PlayListTitle,
				videoStatisticsData.Items[0].Snippet.Title,
				videoId,
				caliData.Season,
				videoStatisticsData,
				videoStatisticsData.Items[0].Snippet.Thumbnails,
				videoStatisticsData.Items[0].Snippet.PublishedAt)
			if err != nil {
				fmt.Printf("[%s]Error: %s", "Calibrate", err.Error())
			}
			videoList = append(videoList, videoItem)
		}
	}
}
