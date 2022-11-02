package matcher

import "strings"

const FREDIT_TEAM_NAME = "Fredit BRION"

type fredit struct{
	Team string
}

func NewFreditMatcher() *fredit {
	return &fredit{Team: FREDIT_TEAM_NAME}
}

func (d *fredit) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *fredit) filter(team string) bool {
	return strings.Contains(team, "fredit") || strings.Contains(team, "프레딧") || strings.Contains(team, "brion") || strings.Contains(team, "브리온")
}