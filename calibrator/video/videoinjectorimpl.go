package calibratorvideo

import (
	"fmt"
	"sort"

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
		for _, videoId := range caliData.VideoIdList {
			videoStatisticsData, err := videoStatisticsService.TempGetVideoStatistics(videoId)
			if err != nil {
				fmt.Println("[Calibrate] Error:", err)
			}
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
			fmt.Printf("[Calibrate]Success append %s %s\n", videoId, date)
		}

		sort.Slice(videoList, func(i, j int) bool {
			return videoList[i].PublishedAt.After(videoList[j].PublishedAt)
		})

		videoMapper[date] = videoList
	}
}
