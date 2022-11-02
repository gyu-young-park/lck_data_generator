package matcher

import "strings"

const LIIV_SANDBOX_TEAM_NAME = "Liiv Sandbox"

type liivSandbox struct{
	Team string
}

func NewLiivSandboxMatcher() *liivSandbox {
	return &liivSandbox{Team: LIIV_SANDBOX_TEAM_NAME}
}

func (d *liivSandbox) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *liivSandbox) filter(team string) bool {
	return strings.Contains(team, "liiv") || strings.Contains(team, "sandbox") || strings.Contains(team, "리브") || strings.Contains(team, "샌드박스")
}