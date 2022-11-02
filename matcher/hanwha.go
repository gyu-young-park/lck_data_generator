package matcher

import "strings"

const HANWHA_TEAM_NAME = "Hanwha Life"

type hanwha struct{
	Team string
}

func NewHanwhaMatcher() *hanwha {
	return &hanwha{Team: HANWHA_TEAM_NAME}
}

func (d *hanwha) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *hanwha) filter(team string) bool {
	return strings.Contains(team, "hanwha") || strings.Contains(team, "한화")
}