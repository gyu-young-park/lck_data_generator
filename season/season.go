package season

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/repository"
	"github.com/gyu-young-park/lck_data_generator/team"
)

type SeasonMapperWithTeam map[string]team.Set

func GenerateSeasonWithTeam(matchList *repository.LCKMatchListModel) SeasonMapperWithTeam {
	if matchList.Error != "null" {
		return nil
	}
	fmt.Println("Start:", "GenerateSeasonWithTeam")
	mapper := SeasonMapperWithTeam{}
	for _, data := range matchList.Data {
		if _, ok := mapper[data.Team1]; !ok {
			mapper[data.Team1] = team.Set{}
		}
		if _, ok := mapper[data.Team2]; !ok {
			mapper[data.Team2] = team.Set{}
		}
		mapper[data.Team1][data.Season] = struct{}{}
		mapper[data.Team2][data.Season] = struct{}{}
	}
	return mapper
}