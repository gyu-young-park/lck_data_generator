package app

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"

	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	"github.com/gyu-young-park/lck_data_generator/repository"
	"github.com/gyu-young-park/lck_data_generator/season"
	"github.com/gyu-young-park/lck_data_generator/team"
)

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func getDateFromVideoTitle(dateParser *regexp.Regexp, videoItem *playlistitems.VideoItemModel) string {
	is19Season := false
	res := string(dateParser.Find([]byte(videoItem.Snippet.Title)))
	if res == "" {
		dateParser, _ = regexp.Compile("(0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01])")
		res = string(dateParser.Find([]byte(videoItem.Snippet.Title)))
		if res == "" {
			fmt.Printf("Can't find data: %s\n", videoItem.Snippet.Title)
		}
		is19Season = true
	}
	if !is19Season {
		res = strings.Split(res, " ")[1]
	}
	fmt.Println(videoItem.Snippet.Title, ": ", res)
	return res
}

func getSeasonFromPlayList(playListItem *playlist.PlaylistItemModel) string {
	var season string
	var buf bytes.Buffer

	playlistSplited := strings.Split(playListItem.Snippet.Title, " ")
	for _, word := range playlistSplited {
		if strings.Contains(word, "게임") || strings.Contains(word, "세트") {
			break
		}
		buf.WriteString(word)
		buf.WriteString(" ")
	}

	season = strings.TrimSpace(buf.String())
	if strings.Contains(season, "[") {
		season = SplitAny(season, "[]")[0]
	}
	return season
}

func makeTeamListWithSeason(matchList *repository.LCKMatchListModel) *repository.LCKTeamWithSeasonListModel {
	teamListWithSeason := repository.LCKTeamWithSeasonListModel{}
	teamMapperWithSeason := team.GenerateTeamWithSeason(matchList)
	if teamMapperWithSeason == nil {
		teamListWithSeason.Error = "Erorr"
		fmt.Println("Error!! team mapper:Can't get team with season")
	} else {
		for season, teamSet := range teamMapperWithSeason {
			var teamWithSeason repository.LCKTeamWithSeasonModel
			teamWithSeason.Season = season
			for team, _ := range teamSet {
				teamWithSeason.TeamList = append(teamWithSeason.TeamList, team)
			}
			teamListWithSeason.Data = append(teamListWithSeason.Data, teamWithSeason)
		}
		teamListWithSeason.Error = "null"
	}
	return &teamListWithSeason
}

func makeSeasonWithTeamList(matchList *repository.LCKMatchListModel) *repository.LCKSeasonWithTeamListModel {
	seasonListWithTeam := repository.LCKSeasonWithTeamListModel{}
	seasonMapperWithTeam := season.GenerateSeasonWithTeam(matchList)
	if seasonMapperWithTeam == nil {
		seasonListWithTeam.Error = "Error"
		fmt.Println("Error!! team mapper:Can't get season with team")
	} else {
		for team, seasonSet := range seasonMapperWithTeam {
			var seasonWithTeam repository.LCKSeasonWithTeamModel
			seasonWithTeam.Team = team
			for season, _ := range seasonSet {
				seasonWithTeam.SeasonList = append(seasonWithTeam.SeasonList, season)
			}
			seasonListWithTeam.Data = append(seasonListWithTeam.Data, seasonWithTeam)
		}
		seasonListWithTeam.Error = "null"
	}
	return &seasonListWithTeam
}
