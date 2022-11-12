package videostatistics

type VideoStatisticsModel struct {
	Views string
}

func NewVideoStatisticsModel(views string) *VideoStatisticsModel {
	return &VideoStatisticsModel{Views: views}
}

type VideoStatisticsResponseModel struct {
	Kind  string `json:"kind"`
	Etag  string `json:"etag"`
	Items []struct {
		Kind       string `json:"kind"`
		Etag       string `json:"etag"`
		ID         string `json:"id"`
		Statistics struct {
			ViewCount     string `json:"viewCount"`
			FavoriteCount string `json:"favoriteCount"`
			CommentCount  string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
	PageInfo struct {
		TotalResults   int `json:"totalResults"`
		ResultsPerPage int `json:"resultsPerPage"`
	} `json:"pageInfo"`
}
