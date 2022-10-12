package playlist

type QueryOption struct{
	Key string
	ChannelId string
	Part string
	Next string
	Max int32 // 0 ~ 50
}

func NewQueryOption(key string, channelId string, part string, next string, max int32) *QueryOption {
	return &QueryOption{key, channelId, part,next, max}
}