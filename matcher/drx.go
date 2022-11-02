package matcher

import "strings"

const DRX_TEAM_NAME = "DRX"

type drx struct{
	Team string
}

func NewDRXMatcher() *drx {
	return &drx{Team: DRX_TEAM_NAME}
}

func (d *drx) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *drx) filter(team string) bool {
	return strings.Contains(team, "king") || strings.Contains(team, "디알") || strings.Contains(team, "dragon") || strings.Contains(team, "drx")
}