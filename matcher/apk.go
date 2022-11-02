package matcher

import "strings"

const SEOL_HAE_ONE_TEAM_NAME = "SeolHaeOne Prince"

type apk struct{
	Team string
}

func NewSeolHaeOneMatcher() *apk {
	return &apk{Team: SEOL_HAE_ONE_TEAM_NAME}
}

func (d *apk) Is(team string) (string, bool){
	if !d.filter(strings.ToLower(team)) {
		return "", false
	}
	return d.Team, true
}

func (d *apk) filter(team string) bool {
	return strings.Contains(team, "apk") || strings.Contains(team, "seol") || strings.Contains(team, "prince") || strings.Contains(team, "설해원") || strings.Contains(team, "프린스")
}