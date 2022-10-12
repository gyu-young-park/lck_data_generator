package playlistitems

type QueryOption struct{
	Key string
	PlaylistId string
	Part string
	Next string
	Max int32 // 0 ~ 50
}

func NewQueryOption(key string, playlistId string, part string, next string, max int32) *QueryOption {
	return &QueryOption{key, playlistId, part,next, max}
}