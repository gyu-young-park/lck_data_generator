package matcher

import "strings"

const NS_TEAM_NAME = "NS RED FORCE"

type ns struct{
	Team string
}

func NewNSMatcher() *ns {
	return &ns{Team: NS_TEAM_NAME}
}

func (d *ns) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *ns) filter(team string) bool {
	return strings.Contains(team, "ns") || strings.Contains(team, "red") || strings.Contains(team, "농심") || strings.Contains(team, "dynamics") || strings.Contains(team, "다이나믹스")
}