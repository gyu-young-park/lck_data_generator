package matcher

import (
	"time"
)

const DEFAULT_DATE_FORMAT = "2006-01-02"

type LCKSeasonAndFandomSeasonMatcher struct {
	fandomMatchRangeList []fandomLCKMatch
}

type fandomLCKMatch struct {
	startDate time.Time
	endDate   time.Time
	season    string
}

func convertTimeWithNoErr(date string) time.Time {
	timeDate, _ := time.Parse(DEFAULT_DATE_FORMAT, date)
	return timeDate
}

func NewLCKSeasonAndFandomSeasonMatcher() *LCKSeasonAndFandomSeasonMatcher {
	return &LCKSeasonAndFandomSeasonMatcher{
		fandomMatchRangeList: []fandomLCKMatch{
			{
				startDate: convertTimeWithNoErr("2019-01-15"),
				endDate:   convertTimeWithNoErr("2019-04-01"),
				season:    "LCK+2019+Spring",
			},
			{
				startDate: convertTimeWithNoErr("2019-04-02"),
				endDate:   convertTimeWithNoErr("2019-04-14"),
				season:    "LCK+2019+Spring+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2019-04-14"),
				endDate:   convertTimeWithNoErr("2019-04-22"),
				season:    "LCK+2019+Summer+Promotion",
			},
			{
				startDate: convertTimeWithNoErr("2019-07-03"),
				endDate:   convertTimeWithNoErr("2019-07-08"),
				season:    "Rift+Rivals+2019+LCK-LPL-LMS-VCS",
			},
			{
				startDate: convertTimeWithNoErr("2019-06-04"),
				endDate:   convertTimeWithNoErr("2019-08-19"),
				season:    "LCK+2019+Summer",
			},
			{
				startDate: convertTimeWithNoErr("2019-08-20"),
				endDate:   convertTimeWithNoErr("2019-09-01"),
				season:    "LCK%2F2019+Season%2FSummer+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2019-09-02"),
				endDate:   convertTimeWithNoErr("2019-09-08"),
				season:    "LCK%2F2019+Season%2FRegional+Finals",
			},
			{
				startDate: convertTimeWithNoErr("2019-09-08"),
				endDate:   convertTimeWithNoErr("2019-09-12"),
				season:    "LCK%2F2020+Season%2FSpring+Promotion",
			},
			{
				startDate: convertTimeWithNoErr("2019-10-01"),
				endDate:   convertTimeWithNoErr("2019-11-11"),
				season:    "2019+Season+World+Championship%2FMain+Event%2C2019+Season+World+Championship%2FPlay-In",
			},
			{
				startDate: convertTimeWithNoErr("2020-02-04"),
				endDate:   convertTimeWithNoErr("2020-04-17"),
				season:    "LCK%2F2020+Season%2FSpring+Season",
			},
			{
				startDate: convertTimeWithNoErr("2020-04-17"),
				endDate:   convertTimeWithNoErr("2020-04-26"),
				season:    "LCK%2F2020+Season%2FSpring+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2020-04-27"),
				endDate:   convertTimeWithNoErr("2020-05-01"),
				season:    "LCK%2F2020+Season%2FSummer+Promotion&_run",
			},
			{
				startDate: convertTimeWithNoErr("2020-05-27"),
				endDate:   convertTimeWithNoErr("2020-06-01"),
				season:    "2020+Mid-Season+Cup",
			},
			{
				startDate: convertTimeWithNoErr("2020-06-16"),
				endDate:   convertTimeWithNoErr("2020-08-24"),
				season:    "LCK%2F2020+Season%2FSummer+Season",
			},
			{
				startDate: convertTimeWithNoErr("2020-08-25"),
				endDate:   convertTimeWithNoErr("2020-09-06"),
				season:    "LCK%2F2020+Season%2FSummer+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2020-09-06"),
				endDate:   convertTimeWithNoErr("2020-09-10"),
				season:    "LCK%2F2020+Season%2FRegional+Finals",
			},
			{
				startDate: convertTimeWithNoErr("2020-09-24"),
				endDate:   convertTimeWithNoErr("2020-11-01"),
				season:    "2020+Season+World+Championship%2FMain+Event%2C2020+Season+World+Championship%2FPlay-In",
			},
			{
				startDate: convertTimeWithNoErr("2021-01-12"),
				endDate:   convertTimeWithNoErr("2021-03-29"),
				season:    "LCK%2F2021+Season%2FSpring+Season",
			},
			{
				startDate: convertTimeWithNoErr("2021-03-30"),
				endDate:   convertTimeWithNoErr("2021-04-11"),
				season:    "LCK%2F2021+Season%2FSpring+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2021-05-05"),
				endDate:   convertTimeWithNoErr("2021-05-24"),
				season:    "2021+Mid-Season+Invitational",
			},
			{
				startDate: convertTimeWithNoErr("2021-06-08"),
				endDate:   convertTimeWithNoErr("2021-08-16"),
				season:    "LCK%2F2021+Season%2FSummer+Season",
			},
			{
				startDate: convertTimeWithNoErr("2021-08-17"),
				endDate:   convertTimeWithNoErr("2021-08-29"),
				season:    "LCK%2F2021+Season%2FSummer+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2021-08-30"),
				endDate:   convertTimeWithNoErr("2021-09-03"),
				season:    "LCK%2F2021+Season%2FRegional+Finals",
			},
			{
				startDate: convertTimeWithNoErr("2021-10-04"),
				endDate:   convertTimeWithNoErr("2021-11-07"),
				season:    "2021+Season+World+Championship%2FMain+Event%2C2021+Season+World+Championship%2FPlay-In",
			},
			{
				startDate: convertTimeWithNoErr("2022-01-11"),
				endDate:   convertTimeWithNoErr("2022-03-21"),
				season:    "LCK%2F2022+Season%2FSpring+Season",
			},
			{
				startDate: convertTimeWithNoErr("2022-03-22"),
				endDate:   convertTimeWithNoErr("2022-04-03"),
				season:    "LCK%2F2022+Season%2FSpring+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2022-05-09"),
				endDate:   convertTimeWithNoErr("2022-05-30"),
				season:    "2022+Mid-Season+Invitational",
			},
			{
				startDate: convertTimeWithNoErr("2022-06-14"),
				endDate:   convertTimeWithNoErr("2022-08-15"),
				season:    "LCK%2F2022+Season%2FSummer+Season",
			},
			{
				startDate: convertTimeWithNoErr("2022-08-16"),
				endDate:   convertTimeWithNoErr("2022-08-29"),
				season:    "LCK%2F2022+Season%2FSummer+Playoffs",
			},
			{
				startDate: convertTimeWithNoErr("2022-08-30"),
				endDate:   convertTimeWithNoErr("2022-09-04"),
				season:    "LCK%2F2022+Season%2FRegional+Finals",
			},
			{
				startDate: convertTimeWithNoErr("2022-09-28"),
				endDate:   convertTimeWithNoErr("2022-11-07"),
				season:    "2022+Season+World+Championship%2FMain+Event%2C2022+Season+World+Championship%2FPlay-In",
			},
		},
	}
}

func (f *fandomLCKMatch) isInTime(date time.Time) bool {
	return date.Before(f.endDate) && date.After(f.startDate)
}

// func (d *LCKSeasonAndFandomSeasonMapper) Is(team string) (string, bool) {
// 	if !d.filter(strings.ToLower(team)) {
// 		return "", false
// 	}
// 	return d.Team, true
// }

func (l *LCKSeasonAndFandomSeasonMatcher) Match(rawDate string) string {
	date := convertTimeWithNoErr(rawDate)
	for _, v := range l.fandomMatchRangeList {
		if v.isInTime(date) {
			return v.season
		}
	}
	return ""
}
