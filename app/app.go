package app

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
	"github.com/gyu-young-park/lck_data_generator/filter"
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
	videoFilter            filter.Filter
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
	app.videoFilter = filter.NewVideoFilter()
	app.FirebaseApp = firebaseapi.NewFireBaseAPI(app.Config.FirebaseKeyPath)
	return app
}

func (app *App) makeLCKPlayList() *[]playlist.PlaylistItemModel {
	channelId, err := app.ChannelService.GetChannelId()
	if err != nil {
		panic(err)
	}
	app.PlayListService = playlist.NewServiceWithChannelId(app.Config.Key, channelId)
	playListItems, err := app.PlayListService.GetPlayListItems()
	if err != nil {
		panic(err)
	}
	return &playListItems
}

func (app *App) makeLCKVideoItems(playListItem *playlist.PlaylistItemModel) *[]playlistitems.VideoItemModel {
	videoItems, err := app.PlayListItemsService.GetVideoItems(playListItem.ID)
	if err != nil {
		panic(err)
	}
	return &videoItems
}

var dateParser, _ = regexp.Compile("\\| (0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01]) \\|")

func (app *App) setVideoItemMapper(videoItemMapper videoitem.VideoItemListMapper, videoItem *playlistitems.VideoItemModel, playListTitle, season string) {
	dateFromTitle := getDateFromVideoTitle(dateParser, videoItem)
	if dateFromTitle == "" {
		fmt.Printf("Error:Date from title is %s\n", dateFromTitle)
		return
	}
	monthDay := strings.Split(dateFromTitle, ".")
	date := fmt.Sprintf("%v-%s-%s", videoItem.Snippet.PublishedAt.Year(), strings.TrimSpace(monthDay[0]), strings.TrimSpace(monthDay[1]))
	// date := videoItem.Snippet.PublishedAt.Format("2006-01-02")
	videostatistics, err := app.VideoStatisticsService.GetVideoStatistics(videoItem.Snippet.ResourceID.VideoID)
	if err != nil {
		fmt.Println(err)
	}
	videoItemMapper.AppendWithDuplicatedCheck(date, videoitem.NewVideoItem(
		playListTitle,
		videoItem.Snippet.Title,
		videoItem.Snippet.ResourceID.VideoID,
		season,
		videostatistics,
		videoItem.Snippet.Thumbnails,
		videoItem.Snippet.PublishedAt))
}

func (app *App) makeLCKVideoItemListMapperWithDate() videoitem.VideoItemListMapper {
	videoItemMapper := make(videoitem.VideoItemListMapper)

	playListItems := app.makeLCKPlayList()
	app.PlayListItemsService = playlistitems.NewServiceWithPlayListId(app.Config.Key)
	for _, playListItem := range *playListItems {
		season := getSeasonFromPlayList(&playListItem)
		videoItems := app.makeLCKVideoItems(&playListItem)
		for _, videoItem := range *videoItems {
			if app.videoFilter.Filtering(videoItem.Snippet.ResourceID.VideoID) {
				fmt.Printf("Filtered: ttile[%s]", videoItem.Snippet.Title)
				continue
			}
			app.setVideoItemMapper(videoItemMapper, &videoItem, playListItem.Snippet.Title, season)
		}
	}
	return videoItemMapper
}

func (app *App) mappingVideoAndResultStage(done chan interface{}, rawResultChan chan interface{}, date string, videoList videoitem.VideoItemList) chan repository.LCKMatchModel {
	matchChan := make(chan repository.LCKMatchModel)
	go func() {
		defer close(matchChan)
		for rawResult := range rawResultChan {
			setResultData := rawResult.([]*crawler.LCKSetDataModel)
			for _, item := range setResultData {
				fmt.Println(item)
			}

			for i, videoItem := range videoList {
				var matchModel repository.LCKMatchModel
				fmt.Println("------------------------")
				matchModel.SetLCKMatchVideo(
					videoItem.PlayList,
					videoItem.Title,
					videoItem.VideoId,
					videoItem.Season,
					videoItem.Statistics.Views,
					videoItem.Thumbnails,
					date,
					videoItem.PublishedAt.Unix(),
				)
				fmt.Println("playlist:", videoItem.PlayList)
				fmt.Println("title:", videoItem.Title)
				fmt.Println("video:", videoItem.VideoId)
				fmt.Println("season:", videoItem.Season)
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
					fmt.Println("date:", videoItem.PublishedAt)
					fmt.Println("------------------------")
				} else {
					matchModel.IsError = true
					fmt.Println("Error ", setResultData, " ", i)
				}
				select {
				case <-done:
					return
				case matchChan <- matchModel:
				}
			}
		}
	}()
	return matchChan
}

func (app *App) makeMatchAndErrorList() (*repository.LCKMatchListModel, *repository.LCKMatchListModel) {
	videoItemMapper := app.makeLCKVideoItemListMapperWithDate()
	matchList := repository.LCKMatchListModel{}
	errorMatchList := repository.LCKMatchListModel{}
	done := make(chan interface{})
	defer close(done)
	// for date, _ := range videoItemMapper {
	// 	app.crawler.SetData(date)
	// 	dataSetResultMapper[date] = app.crawler.GoroutineGetResult(done)
	// }
	// for rawSetResultData := range dataSetResultMapper[date].(<-chan interface{}) {
	// 	rawSetResultData.([]*crawler.LCKSetDataModel)
	// }
	mapperMatchWithDate := make(map[string]chan repository.LCKMatchModel)
	// matchChanList := make([]chan repository.LCKMatchModel, 10)
	for date, videoList := range videoItemMapper {
		app.crawler.SetData(date)
		mapperMatchWithDate[date] = app.mappingVideoAndResultStage(done, app.crawler.GoroutineGetResult(done), date, videoList)
	}
	for _, matchChan := range mapperMatchWithDate {
		for match := range matchChan {
			if match.IsError {
				errorMatchList.Data = append(errorMatchList.Data, match)
			} else {
				matchList.Data = append(matchList.Data, match)
			}
			close(matchChan)
		}
	}
	return &matchList, &errorMatchList
}

func (app *App) removeAllDBSchema() {
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
	err = app.FirebaseApp.RemoveCollection("lck_teams")
	if err != nil {
		fmt.Println(err)
	}
	err = app.FirebaseApp.RemoveCollection("lck_seasons")
	if err != nil {
		fmt.Println(err)
	}
}

func (app *App) storeAllDataInFirebase(
	matchList *repository.LCKMatchListModel,
	teamListWithSeason *repository.LCKTeamWithSeasonListModel,
	seasonListWithTeam *repository.LCKSeasonWithTeamListModel,
	teamList *team.TeamList,
	seasonList *season.SeasonList,
) {
	for _, matchData := range matchList.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_match", matchData.VideoId, firebaseapi.FireStoreDataSchema(structs.Map(matchData)))
	}

	for _, teamWithSeasonData := range teamListWithSeason.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_team_with_season", teamWithSeasonData.Season, firebaseapi.FireStoreDataSchema(structs.Map(teamWithSeasonData)))
	}

	for _, seasonWithTeamData := range seasonListWithTeam.Data {
		app.FirebaseApp.StoreDataWithDoc("lck_season_with_team", seasonWithTeamData.Team, firebaseapi.FireStoreDataSchema(structs.Map(seasonWithTeamData)))
	}
	app.FirebaseApp.StoreDataWithDoc("lck_teams", "teams", firebaseapi.FireStoreDataSchema(structs.Map(teamList)))
	app.FirebaseApp.StoreDataWithDoc("lck_seasons", "seasons", firebaseapi.FireStoreDataSchema(structs.Map(seasonList)))
	fmt.Println("teams:", teamList)
	fmt.Println("seasons", seasonList)
}

func (app *App) storeAllDataInJSONFile(
	matchList *repository.LCKMatchListModel,
	errorMatchList *repository.LCKMatchListModel,
	teamListWithSeason *repository.LCKTeamWithSeasonListModel,
	seasonListWithTeam *repository.LCKSeasonWithTeamListModel,
	teamList *team.TeamList,
	seasonList *season.SeasonList,
) {
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
	data, err = json.MarshalIndent(teamList, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.Repo.Store(string(repository.ALL_TEAM_LIST), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err = json.MarshalIndent(seasonList, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.Repo.Store(string(repository.ALL_SEASON_LIST), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err = json.MarshalIndent(errorMatchList, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = app.Repo.Store(string(repository.ALL_ERROR_MATCH_LIST), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (app *App) Start() {
	matchList, errorMatchList := app.makeMatchAndErrorList()
	if len(matchList.Data) == 0 {
		matchList.Error = "Error: There are no data"
		return
	}
	matchList.Error = "null"
	teamListWithSeason := makeTeamListWithSeason(matchList)
	seasonListWithTeam := makeSeasonWithTeamList(matchList)
	teamList := team.GenerateTeamList()
	seasonList := season.GenerateSeasonList(matchList)

	app.removeAllDBSchema()
	app.storeAllDataInFirebase(matchList, teamListWithSeason, seasonListWithTeam, teamList, seasonList)
	app.storeAllDataInJSONFile(matchList, errorMatchList, teamListWithSeason, seasonListWithTeam, teamList, seasonList)
	app.server.StartServer()
}

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
