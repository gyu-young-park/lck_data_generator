package matcher

import "strings"

const KWANGDONG_TEAM_NAME = "Kwangdong Freecs"

type kwangdong struct{
	Team string
}

func NewKwangdongMatcher() *kwangdong {
	return &kwangdong{Team: KWANGDONG_TEAM_NAME}
}

func (d *kwangdong) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *kwangdong) filter(team string) bool {
	return strings.Contains(team, "kwangdong") || strings.Contains(team, "freecs") || strings.Contains(team, "afreeca") || strings.Contains(team, "프릭스") || strings.Contains(team, "아프리카") || strings.Contains(team, "광동")
}