package matcher

import "strings"

const GRIFFIN_TEAM_NAME = "Griffin"

type griffin struct{
	Team string
}

func NewGriffinMatcher() *griffin {
	return &griffin{Team: GRIFFIN_TEAM_NAME}
}

func (d *griffin) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *griffin) filter(team string) bool {
	return strings.Contains(team, "griffin") || strings.Contains(team, "그리핀")
}