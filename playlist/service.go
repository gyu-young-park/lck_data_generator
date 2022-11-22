package playlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gyu-young-park/lck_data_generator/filter"
)

type Service interface {
	GetPlayListItems(filterOpt filter.Filter) ([]PlaylistItemModel, error)
}

const PART_OPTION = "snippet"

type ServiceWithChannelId struct {
	option *QueryOption
}

func NewServiceWithChannelId(key string, channelId string) *ServiceWithChannelId {
	ins := &ServiceWithChannelId{NewQueryOption(key, channelId, PART_OPTION, "", 50)}
	return ins
}

func (s *ServiceWithChannelId) GetPlayListItems(filterOpt filter.Filter) ([]PlaylistItemModel, error) {
	var playListItems []PlaylistItemModel
	for {
		url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?channelId=%s&part=%s&key=%s&maxResults=%d&pageToken=%s",
			s.option.ChannelId, s.option.Part, s.option.Key, s.option.Max, s.option.Next)
		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		var playlist PlaylistModel
		err = json.Unmarshal(data, &playlist)
		if err != nil {
			panic(err)
		}
		for _, item := range playlist.Items {
			if item.ID != "" {
				if filterOpt.Filtering(item.Snippet.Title) {
					playListItems = append(playListItems, item)
				}
			}
		}
		s.option.Next = playlist.NextPageToken
		if playlist.NextPageToken == "" {
			return playListItems, nil
		}
	}
}
