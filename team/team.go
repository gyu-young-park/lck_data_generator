package team

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/commontype"
	"github.com/gyu-young-park/lck_data_generator/matcher"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

type TeamMapperWithSeason map[string]commontype.Set

func GenerateTeamWithSeason(matchList *repository.LCKMatchListModel) TeamMapperWithSeason {
	if matchList.Error != "null" {
		return nil
	}
	fmt.Println("Start:", "GenerateTeamWithSeason")
	mapper := TeamMapperWithSeason{}
	for _, data := range matchList.Data {
		if _, ok := mapper[data.Season]; !ok {
			mapper[data.Season] = commontype.Set{}
		}
		mapper[data.Season][data.Team1] = struct{}{}
		mapper[data.Season][data.Team2] = struct{}{}
	}
	return mapper
}

type TeamList struct {
	Teams []string `json:"teams"`
}

var teamList = []string{
	matcher.DK_TEAM_NAME,
	matcher.DRX_TEAM_NAME,
	matcher.FREDIT_TEAM_NAME,
	matcher.GENG_TEAM_NAME,
	matcher.GRIFFIN_TEAM_NAME,
	matcher.HANWHA_TEAM_NAME,
	matcher.JINAIR_TEAM_NAME,
	matcher.KT_TEAM_NAME,
	matcher.KWANGDONG_TEAM_NAME,
	matcher.LIIV_SANDBOX_TEAM_NAME,
	matcher.NS_TEAM_NAME,
	matcher.SEOL_HAE_ONE_TEAM_NAME,
	matcher.T1_TEAM_NAME,
}

func GenerateTeamList() *TeamList {
	fmt.Println("Start:", "GenerateTeamList")
	var ret TeamList
	for _, team := range teamList {
		ret.Teams = append(ret.Teams, team)
	}
	return &ret
}
