package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/fatih/structs"
	"github.com/gyu-young-park/lck_data_generator/api"
	"github.com/gyu-young-park/lck_data_generator/channel"
	"github.com/gyu-young-park/lck_data_generator/config"
	"github.com/gyu-young-park/lck_data_generator/crawler"
	"github.com/gyu-young-park/lck_data_generator/firebaseapi"
	"github.com/gyu-young-park/lck_data_generator/matcher"
	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	"github.com/gyu-young-park/lck_data_generator/repository"
	"github.com/gyu-young-park/lck_data_generator/season"
	"github.com/gyu-young-park/lck_data_generator/team"
	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
	"github.com/gyu-young-park/lck_data_generator/videostatistics"
)

type App struct {
	Config                 *config.Config
	server                 *api.Server
	teamMatcher            matcher.Matcher
	crawler                crawler.Crawler
	ChannelService         channel.Service
	PlayListService        playlist.Service
	PlayListItemsService   playlistitems.Service
	VideoStatisticsService videostatistics.Service
	FirebaseApp            *firebaseapi.FirebaseApp
	Repo                   repository.Repository
}

func NewApp() *App {
	app := &App{crawler: crawler.NewLCKSetResultCrawler()}
	app.teamMatcher = matcher.NewLCKTeamMatcher()
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
	app.VideoStatisticsService = videostatistics.NewServiceWithVideoStatistics(app.Config.Key)
	app.Repo = repository.NewFileRepository()
	app.server = api.NewHTTPServer()
	app.FirebaseApp = firebaseapi.NewFireBaseAPI(app.Config.FirebaseKeyPath)
	return app
}

func SplitAny(s string, seps string) []string {
	splitter := func(r rune) bool {
		return strings.ContainsRune(seps, r)
	}
	return strings.FieldsFunc(s, splitter)
}

func (app *App) MakeLCKVideoItemList() videoitem.VideoItemListMapper {
	videoItemMapper := make(videoitem.VideoItemListMapper)
	channelId, err := app.ChannelService.GetChannelId()
	if err != nil {
		panic(err)
	}
	app.PlayListService = playlist.NewServiceWithChannelId(app.Config.Key, channelId)
	playListItems, err := app.PlayListService.GetPlayListItems()
	if err != nil {
		panic(err)
	}
	app.PlayListItemsService = playlistitems.NewServiceWithPlayListId(app.Config.Key)
	for _, playListItem := range playListItems {
		videoItems, err := app.PlayListItemsService.GetVideoItems(playListItem.ID)
		if err != nil {
			panic(err)
		}
		for _, videoItem := range videoItems {
			dateParser, _ := regexp.Compile("\\| (0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01]) \\|")
			is19Season := false
			res := string(dateParser.Find([]byte(videoItem.Snippet.Title)))
			if res == "" {
				dateParser, _ = regexp.Compile("(0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01])")
				res = string(dateParser.Find([]byte(videoItem.Snippet.Title)))
				if res == "" {
					fmt.Printf("Can't find data: %s\n", videoItem.Snippet.Title)
					continue
				}
				is19Season = true
			}
			if !is19Season {
				res = strings.Split(res, " ")[1]
			}
			fmt.Println(videoItem.Snippet.Title, ": ", res)
			monthDay := strings.Split(res, ".")
			if len(monthDay) != 2 {
				fmt.Printf("debug monthDay:%v\n", monthDay)
				continue
			}
			// titles := SplitAny(videoItem.Snippet.Title, "|[]")
			playlistSplited := strings.Split(playListItem.Snippet.Title, " ")
			var buf bytes.Buffer
			for _, word := range playlistSplited {
				if strings.Contains(word, "게임") || strings.Contains(word, "세트") {
					break
				}
				buf.WriteString(word)
				buf.WriteString(" ")
			}
			var season string
			// if len(titles) == 0 || len(titles) == 1 {
			// 	// titles := strings.Split(videoItem.Snippet.Title, ` `)
			// 	// fmt.Println(titles)
			// 	// if len(titles) == 0 || len(titles) == 1 {
			// 	// 	fmt.Printf("Season Error:%v\n", titles)
			// 	// 	season = "null"
			// 	// } else {
			// 	// 	season = titles[len(titles) - 1]
			// 	// }
			// 	fmt.Printf("Season Error:%v\n", titles)
			// 		season = "null"
			// } else {
			// 	season = titles[len(titles) - 1]
			// }
			season = strings.TrimSpace(buf.String())
			if strings.Contains(season, "[") {
				season = SplitAny(season, "[]")[0]
			}
			date := fmt.Sprintf("%v-%s-%s", videoItem.Snippet.PublishedAt.Year(), strings.TrimSpace(monthDay[0]), strings.TrimSpace(monthDay[1]))
			// date := videoItem.Snippet.PublishedAt.Format("2006-01-02")
			videostatistics, err := app.VideoStatisticsService.GetVideoStatistics(videoItem.Snippet.ResourceID.VideoID)
			if err != nil {
				fmt.Println(err)
			}
			videoItemMapper[date] = append(videoItemMapper[date], videoitem.NewVideoItem(
				playListItem.Snippet.Title,
				videoItem.Snippet.Title,
				videoItem.Snippet.ResourceID.VideoID,
				season,
				videostatistics,
				videoItem.Snippet.Thumbnails,
				videoItem.Snippet.PublishedAt))
		}
	}
	return videoItemMapper
}

func main() {
	app := NewApp()
	videoItemMapper := app.MakeLCKVideoItemList()
	matchList := repository.LCKMatchListModel{}
	teamListWithSeason := repository.LCKTeamWithSeasonListModel{}
	seasonListWithTeam := repository.LCKSeasonWithTeamListModel{}
	for k, v := range videoItemMapper {
		app.crawler.SetData(k)
		rawSetResultData := app.crawler.GetResult()
		setResultData, _ := rawSetResultData.([]*crawler.LCKSetDataModel)
		for _, item := range setResultData {
			fmt.Println(item)
		}

		for i, item := range v {
			var matchModel repository.LCKMatchModel
			fmt.Println("------------------------")
			matchModel.SetLCKMatchVideo(
				item.PlayList,
				item.Title,
				item.VideoId,
				item.Season,
				item.Statistics.Views,
				item.Thumbnails,
				k,
				item.PublishedAt.Unix(),
			)
			fmt.Println("playlist:", item.PlayList)
			fmt.Println("title:", item.Title)
			fmt.Println("video:", item.VideoId)
			fmt.Println("season:", item.Season)
			if len(setResultData) > i {
				team1 := app.teamMatcher.Match(setResultData[i].TeamScore1.Team)
				team2 := app.teamMatcher.Match(setResultData[i].TeamScore2.Team)
				matchModel.SetLCKMatchScore(
					team1,
					setResultData[i].TeamScore1.Score,
					team2,
					setResultData[i].TeamScore2.Score)
				fmt.Println("team1:", team1)
				fmt.Println("team1-result:", setResultData[i].TeamScore1.Score)
				fmt.Println("team2:", team2)
				fmt.Println("team2-result:", setResultData[i].TeamScore2.Score)
			} else {
				matchModel.IsError = true
				fmt.Println("Error ", setResultData, " ", i)
			}
			fmt.Println("date:", item.PublishedAt)
			fmt.Println("------------------------")
			matchList.Data = append(matchList.Data, matchModel)
		}
	}

	if len(matchList.Data) == 0 {
		matchList.Error = "Error: There are no data"
	}
	matchList.Error = "null"
	teamMapperWithSeason := team.GenerateTeamWithSeason(&matchList)
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

	seasonMapperWithTeam := season.GenerateSeasonWithTeam(&matchList)
	if seasonMapperWithTeam == nil {
		seasonListWithTeam.Error = "Error"
		fmt.Println("Error!! team mapper:Can't get season with team")
	} else {
		for team, seasonSet := range teamMapperWithSeason {
			var seasonWithTeam repository.LCKSeasonWithTeamModel
			seasonWithTeam.Team = team
			for season, _ := range seasonSet {
				seasonWithTeam.SeasonList = append(seasonWithTeam.SeasonList, season)
			}
			seasonListWithTeam.Data = append(seasonListWithTeam.Data, seasonWithTeam)
		}
		seasonListWithTeam.Error = "null"
	}

	app.FirebaseApp.StoreDump()
	dumpList := app.FirebaseApp.ReadDump()
	fmt.Println("dump start-----")
	for _, fireSotreSchema := range dumpList {
		fmt.Println("-----step----")
		for k, v := range fireSotreSchema {
			fmt.Printf("key:%s value:%v\n", k, v)
		}
	}
	err := app.FirebaseApp.RemoveCollection("lck_match")
	if err != nil {
		fmt.Println(err)
	}
	err = app.FirebaseApp.RemoveCollection("lck_season_with_team")
	if err != nil {
		fmt.Println(err)
	}
	err = app.FirebaseApp.RemoveCollection("lck_team_with_season")
	if err != nil {
		fmt.Println(err)
	}
	for _, matchData := range matchList.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_match", matchData.VideoId, firebaseapi.FireStoreDataSchema(structs.Map(matchData)))
	}

	for _, teamWithSeasonData := range teamListWithSeason.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_season_with_team", teamWithSeasonData.Season, firebaseapi.FireStoreDataSchema(structs.Map(teamWithSeasonData)))
	}

	for _, seasonTeamData := range seasonListWithTeam.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_team_with_season", seasonTeamData.Team, firebaseapi.FireStoreDataSchema(structs.Map(seasonTeamData)))
	}

	data, err := json.MarshalIndent(matchList, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = app.Repo.Store(string(repository.ALL_MATCH), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = json.MarshalIndent(teamListWithSeason, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = app.Repo.Store(string(repository.ALL_TEAM_WITH_SEASON), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = json.MarshalIndent(seasonListWithTeam, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = app.Repo.Store(string(repository.ALL_SEASON_WITH_TEAM), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	app.server.StartServer()
}
