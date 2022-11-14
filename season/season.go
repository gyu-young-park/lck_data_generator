package season

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/commontype"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

type SeasonMapperWithTeam map[string]commontype.Set

func GenerateSeasonWithTeam(matchList *repository.LCKMatchListModel) SeasonMapperWithTeam {
	if matchList.Error != "null" {
		return nil
	}
	fmt.Println("Start:", "GenerateSeasonWithTeam")
	mapper := SeasonMapperWithTeam{}
	for _, data := range matchList.Data {
		if _, ok := mapper[data.Team1]; !ok {
			mapper[data.Team1] = commontype.Set{}
		}
		if _, ok := mapper[data.Team2]; !ok {
			mapper[data.Team2] = commontype.Set{}
		}
		mapper[data.Team1][data.Season] = struct{}{}
		mapper[data.Team2][data.Season] = struct{}{}
	}
	return mapper
}

type SeasonList struct {
	SeasonList []string `json:"seasons"`
}

func GenerateSeasonList(matchList *repository.LCKMatchListModel) *SeasonList {
	if matchList.Error != "null" {
		return nil
	}
	fmt.Println("Start:", "GenerateSeasonList")
	var ret SeasonList
	seasonSet := commontype.Set{}
	for _, data := range matchList.Data {
		seasonSet[data.Season] = struct{}{}
	}
	for season, _ := range seasonSet {
		ret.SeasonList = append(ret.SeasonList, season)
	}
	return &ret
}
