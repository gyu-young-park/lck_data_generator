package matcher

import "strings"

const T1_TEAM_NAME = "T1"

type t1 struct{
	Team string
}

func NewT1Matcher() *t1 {
	return &t1{Team: T1_TEAM_NAME}
}

func (d *t1) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *t1) filter(team string) bool {
	return strings.Contains(team, "t1") || strings.Contains(team, "sk") || strings.Contains(team, "telecom") || strings.Contains(team, "티원") || strings.Contains(team, "에스케이")
}