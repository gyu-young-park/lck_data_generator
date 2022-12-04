package videostatistics

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

type TempViewStruct struct {
	Kind    string "json:\"kind\""
	Etag    string "json:\"etag\""
	ID      string "json:\"id\""
	Snippet struct {
		PublishedAt time.Time "json:\"publishedAt\""
		ChannelID   string    "json:\"channelId\""
		Title       string    "json:\"title\""
		Description string    "json:\"description\""
		Thumbnails  struct {
			Default struct {
				URL    string "json:\"url\""
				Width  int    "json:\"width\""
				Height int    "json:\"height\""
			} "json:\"default\""
			Medium struct {
				URL    string "json:\"url\""
				Width  int    "json:\"width\""
				Height int    "json:\"height\""
			} "json:\"medium\""
			High struct {
				URL    string "json:\"url\""
				Width  int    "json:\"width\""
				Height int    "json:\"height\""
			} "json:\"high\""
			Standard struct {
				URL    string "json:\"url\""
				Width  int    "json:\"width\""
				Height int    "json:\"height\""
			} "json:\"standard\""
			Maxres struct {
				URL    string "json:\"url\""
				Width  int    "json:\"width\""
				Height int    "json:\"height\""
			} "json:\"maxres\""
		} "json:\"thumbnails\""
		ChannelTitle         string   "json:\"channelTitle\""
		Tags                 []string "json:\"tags\""
		CategoryID           string   "json:\"categoryId\""
		LiveBroadcastContent string   "json:\"liveBroadcastContent\""
		DefaultLanguage      string   "json:\"defaultLanguage\""
		Localized            struct {
			Title       string "json:\"title\""
			Description string "json:\"description\""
		} "json:\"localized\""
		DefaultAudioLanguage string "json:\"defaultAudioLanguage\""
	} "json:\"snippet\""
	Statistics struct {
		ViewCount     string "json:\"viewCount\""
		FavoriteCount string "json:\"favoriteCount\""
		CommentCount  string "json:\"commentCount\""
	} "json:\"statistics\""
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
	// err = json.Unmarshal(data, &videoStatistics)
	// if err != nil {
	// 	panic(err)
	// }
	// if len(videoStatistics.Items) != 1 {
	// 	return videoStatistics, fmt.Errorf("Error views are not valid%v\n", len(videoStatistics.Items))
	// }
	videoStatistics.Items = append(videoStatistics.Items, TempViewStruct{})
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
	err = json.Unmarshal(data, &videoStatistics)
	if err != nil {
		panic(err)
	}
	if len(videoStatistics.Items) != 1 {
		return videoStatistics, fmt.Errorf("Error views are not valid%v\n", len(videoStatistics.Items))
	}
	return videoStatistics, nil
}
