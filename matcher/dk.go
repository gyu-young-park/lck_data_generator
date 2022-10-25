package matcher

import "strings"

const DK_TEAM_NAME = "DK"

type dk struct{
	Team string
}

func NewDKMatcher() *dk {
	return &dk{Team: DK_TEAM_NAME}
}

func (d *dk) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *dk) filter(team string) bool {
	return strings.Contains(team, "dk") || strings.Contains(team, "dwg") || strings.Contains(team, "담원") || strings.Contains(team, "damwon") || strings.Contains(team, "디케이")
}