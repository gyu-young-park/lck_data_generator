package team

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/repository"
)

type Set map[string]struct{}

type TeamMapperWithSeason map[string]Set

func GenerateTeamWithSeason(matchList *repository.LCKMatchListModel) TeamMapperWithSeason {
	if matchList.Error != "null" {
		return nil
	}
	fmt.Println("Start:", "GenerateTeamWithSeason")
	mapper := TeamMapperWithSeason{}
	for _, data := range matchList.Data {
		if _, ok := mapper[data.Season]; !ok {
			mapper[data.Season] = Set{}
		}
		mapper[data.Season][data.Team1] = struct{}{}
		mapper[data.Season][data.Team2] = struct{}{}
	}
	return mapper
}