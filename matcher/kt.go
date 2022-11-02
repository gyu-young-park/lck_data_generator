package matcher

import "strings"

const KT_TEAM_NAME = "KT Rolster"

type kt struct{
	Team string
}

func NewKTMatcher() *kt {
	return &kt{Team: KT_TEAM_NAME}
}

func (d *kt) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *kt) filter(team string) bool {
	return strings.Contains(team, "kt") || strings.Contains(team, "rolster") || strings.Contains(team, "케이티")
}