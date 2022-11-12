package videostatistics

type QueryOption struct {
	Key     string
	VideoId string
	Part    string
}

func NewQueryOption(key string, videoId string, part string) *QueryOption {
	return &QueryOption{key, videoId, part}
}
