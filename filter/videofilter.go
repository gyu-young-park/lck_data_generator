package filter

var EXCEPT_VIDEO_MAP = []string{
	"l-4eP1z7huE",
}

type videoFilterExceptMapper map[string]bool
type VideoFilter struct {
	exceptVideoMapper videoFilterExceptMapper
}

func NewVideoFilter() *VideoFilter {
	videoFilter := &VideoFilter{
		exceptVideoMapper: make(videoFilterExceptMapper),
	}
	videoFilter.setUp()
	return videoFilter
}

func (v *VideoFilter) setUp() {
	for _, videoId := range EXCEPT_VIDEO_MAP {
		v.exceptVideoMapper[videoId] = true
	}
}

func (v *VideoFilter) Filtering(data interface{}) bool {
	videoId, ok := data.(string)
	if !ok {
		return false
	}
	return v.exceptVideoMapper[videoId]
}
