package playlistitems

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Service interface{
	GetVideoItems(playlistId string) ([]VideoItemModel, error)
}

const PART_OPTION ="snippet"

type ServiceWithPlayListId struct {
	option *QueryOption
}
	
func NewServiceWithPlayListId(key string) *ServiceWithPlayListId {
	ins := &ServiceWithPlayListId{NewQueryOption(key, "", PART_OPTION, "", 50)}
	return ins
}

func (s *ServiceWithPlayListId) GetVideoItems(playlistId string) ([]VideoItemModel, error) {
	var videoItemList []VideoItemModel
	s.option.PlaylistId = playlistId
	for {
		url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/playlistItems?playlistId=%s&key=%s&maxResults=%d&pageToken=%s&part=%s", 
						s.option.PlaylistId,s.option.Key,s.option.Max,s.option.Next, s.option.Part)
		res ,err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
		data ,err := ioutil.ReadAll(res.Body)
		var playlistItems PlaylistItemsModel
		err = json.Unmarshal(data, &playlistItems)
		if err != nil {
			panic(err)
		}
		for _, item := range playlistItems.Items {
			if item.Snippet.ChannelID != "" {
				videoItemList = append(videoItemList, item)
			}
		}
		s.option.Next = playlistItems.NextPageToken
		if s.option.Next == "" {
			return videoItemList, nil
		}
	}	
}