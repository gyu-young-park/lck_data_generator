package matcher

type teamMatcher interface{
	Is(string) (string, bool)
}

type LCKTeamMatcher struct{
	teamMatcher []teamMatcher
}

func NewLCKTeamMatcher() *LCKTeamMatcher {
	return &LCKTeamMatcher{
		teamMatcher: []teamMatcher{
			NewDKMatcher(),
		},
	}
}

func (l*LCKTeamMatcher)Match(team string) string {
	for _, v := range l.teamMatcher {
		teamName, ok := v.Is(team)
		if ok {
			return teamName
		}
	}
	return team
}