package playlistitems

import "time"

type PlaylistItemsModel struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	NextPageToken string `json:"nextPageToken"`
	Items []VideoItemModel `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}

type VideoItemModel struct {
	Kind    string `json:"kind"`
	Etag    string `json:"etag"`
	ID      string `json:"id"`
	Snippet VideoSnippetModel `json:"snippet"`
}

type VideoSnippetModel struct {
	PublishedAt time.Time `json:"publishedAt"`
	ChannelID   string    `json:"channelId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnails  struct {
		Default struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"default"`
		Medium struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"medium"`
		High struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"high"`
		Standard struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"standard"`
		Maxres struct {
			URL    string `json:"url"`
			Width  int    `json:"width"`
			Height int    `json:"height"`
		} `json:"maxres"`
	} `json:"thumbnails"`
	ChannelTitle string `json:"channelTitle"`
	PlaylistID   string `json:"playlistId"`
	Position     int    `json:"position"`
	ResourceID   struct {
		Kind    string `json:"kind"`
		VideoID string `json:"videoId"`
	} `json:"resourceId"`
	VideoOwnerChannelTitle string `json:"videoOwnerChannelTitle"`
	VideoOwnerChannelID    string `json:"videoOwnerChannelId"`
}