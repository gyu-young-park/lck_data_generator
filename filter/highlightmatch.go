package filter

import "strings"

type HighlightMatchFilter struct {
	
}

func NewHighlightMatchFilter() *HighlightMatchFilter{
	highlightMatchFilter := &HighlightMatchFilter{}
	return highlightMatchFilter
}

func (h *HighlightMatchFilter)Filtering(data interface{}) bool{
	matchName, ok := data.(string)
	if !ok {
		return false
	}
	lowerMatchName := strings.ToLower(matchName)
	if strings.Contains(lowerMatchName, "match") || strings.Contains(lowerMatchName, "매치") {
		return true
	}
	return false
}