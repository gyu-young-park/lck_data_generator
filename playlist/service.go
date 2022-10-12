package playlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface{
	GetPlayListId() ([]string, error)
}

const PART_OPTION ="snippet"

type ServiceWithChannelId struct {
	option *QueryOption
}
	
func NewServiceWithChannelId(key string, channelId string) *ServiceWithChannelId {
	ins := &ServiceWithChannelId{NewQueryOption(key, channelId, PART_OPTION, "", 50)}
	return ins
}

func (s *ServiceWithChannelId) GetPlayListId() ([]string, error) {
	var playlistIds []string
	for {
		url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlists?channelId=%s&part=%s&key=%s&maxResults=%d&pageToken=%s", 
						s.option.ChannelId,s.option.Part,s.option.Key,s.option.Max, s.option.Next)
		res ,err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		data ,err := ioutil.ReadAll(res.Body)
		var playlist PlaylistModel
		err = json.Unmarshal(data, &playlist)
		if err != nil {
			panic(err)
		}
		for _, item := range playlist.Items {
			if item.ID != "" {
				playlistIds = append(playlistIds, item.ID)
			}
		}
		s.option.Next = playlist.NextPageToken
		if playlist.NextPageToken == "" {
			return playlistIds, nil
		}
	}	
}