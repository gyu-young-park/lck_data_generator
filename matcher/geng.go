package matcher

import "strings"

const GENG_TEAM_NAME = "GEN.G"

type geng struct{
	Team string
}

func NewGenGMatcher() *geng {
	return &geng{Team: GENG_TEAM_NAME}
}

func (d *geng) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *geng) filter(team string) bool {
	return strings.Contains(team, "gen") || strings.Contains(team, "젠지")
}