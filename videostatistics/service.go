package videostatistics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface {
	GetVideoStatistics(videoId string) (VideoStatisticsResponseModel, error)
	TempGetVideoStatistics(videoId string) (VideoStatisticsResponseModel, error)
}

const PART_OPTION = "snippet,statistics"

type ServiceWithVideoStatistics struct {
	option *QueryOption
}

func NewServiceWithVideoStatistics(key string) *ServiceWithVideoStatistics {
	ins := &ServiceWithVideoStatistics{NewQueryOption(key, "", PART_OPTION)}
	return ins
}

func (s *ServiceWithVideoStatistics) GetVideoStatistics(videoId string) (VideoStatisticsResponseModel, error) {
	var videoStatistics VideoStatisticsResponseModel
	// if videoId == "" {
	// 	return videoStatistics, fmt.Errorf("Error videoId is not valid%v\n", videoId)
	// }
	// s.option.VideoId = videoId
	// //url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?channelId=%s&part=%s&key=%s&maxResults=%d&pageToken=%s"
	// url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=%s",
	// 	s.option.VideoId, s.option.Key, s.option.Part)
	// res, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }
	// defer res.Body.Close()
	// data, err := ioutil.ReadAll(res.Body)
	// var videoStatisticsResponse VideoStatisticsResponseModel
	// err = json.Unmarshal(data, &videoStatisticsResponse)
	// if err != nil {
	// 	panic(err)
	// }
	// if len(videoStatisticsResponse.Items) != 1 {
	// 	return videoStatistics, fmt.Errorf("Error views are not valid%v\n", len(videoStatisticsResponse.Items))
	// }
	return videoStatistics, nil
}

func (s *ServiceWithVideoStatistics) TempGetVideoStatistics(videoId string) (VideoStatisticsResponseModel, error) {
	var videoStatistics VideoStatisticsResponseModel
	if videoId == "" {
		return videoStatistics, fmt.Errorf("Error videoId is not valid%v\n", videoId)
	}
	s.option.VideoId = videoId
	//url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?channelId=%s&part=%s&key=%s&maxResults=%d&pageToken=%s"
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&key=%s&part=%s",
		s.option.VideoId, s.option.Key, s.option.Part)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	var videoStatisticsResponse VideoStatisticsResponseModel
	err = json.Unmarshal(data, &videoStatisticsResponse)
	if err != nil {
		panic(err)
	}
	if len(videoStatisticsResponse.Items) != 1 {
		return videoStatistics, fmt.Errorf("Error views are not valid%v\n", len(videoStatisticsResponse.Items))
	}
	return videoStatistics, nil
}
