package matcher

import "strings"

const JINAIR_TEAM_NAME = "Jin Air"

type jinair struct{
	Team string
}

func NewJinairMatcher() *jinair {
	return &jinair{Team: JINAIR_TEAM_NAME}
}

func (d *jinair) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *jinair) filter(team string) bool {
	return strings.Contains(team, "jin") || strings.Contains(team, "air") || strings.Contains(team, "진에어")
}