package filter

import "strings"

type SetHightlightFilter struct {
}

func NewSetHightlightFilter() *SetHightlightFilter {
	setHightlightFilter := &SetHightlightFilter{}
	return setHightlightFilter
}

func (h *SetHightlightFilter) Filtering(data interface{}) bool {
	matchName, ok := data.(string)
	if !ok {
		return false
	}
	lowerMatchName := strings.ToLower(matchName)
	if strings.Contains(lowerMatchName, "set") || (strings.Contains(lowerMatchName, "게임") && strings.Contains(lowerMatchName, "하이라이트")) || strings.Contains(lowerMatchName, "세트") {
		return true
	}
	return false
}
