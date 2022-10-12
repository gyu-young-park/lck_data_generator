package channel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface{
	GetChannelId() (string, error)
}

const SPECIFIC_LCK_VIDEO_ID = "jxr_NONF74k"

type ServiceWithVideoId struct {
	key string
	videoId string
}

func NewServiceWithVideoId(key string) *ServiceWithVideoId {
	ins := &ServiceWithVideoId{key: key, videoId: SPECIFIC_LCK_VIDEO_ID}
	return ins
}

func (s *ServiceWithVideoId) GetChannelId() (string, error) {
	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?id=%s&part=%s&key=%s", s.videoId, "snippet",s.key)
	res ,err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	data ,err := ioutil.ReadAll(res.Body)
	var video Video
	err = json.Unmarshal(data, &video)
	if err != nil {
		panic(err)
	}
	for _, item := range video.Items {
		if item.ID != "" {
			return item.ID, nil
		}
	}
	return "", fmt.Errorf("ERR!")
}