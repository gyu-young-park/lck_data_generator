package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/gyu-young-park/lck_data_generator/api"
	"github.com/gyu-young-park/lck_data_generator/channel"
	"github.com/gyu-young-park/lck_data_generator/config"
	"github.com/gyu-young-park/lck_data_generator/crawler"
	"github.com/gyu-young-park/lck_data_generator/matcher"
	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	"github.com/gyu-young-park/lck_data_generator/repository"
	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
)

type App struct {
	Config               *config.Config
	server               *api.Server
	teamMatcher          matcher.Matcher
	crawler              crawler.Crawler
	ChannelService       channel.Service
	PlayListService      playlist.Service
	PlayListItemsService playlistitems.Service
	Repo                 repository.Repository
}

func NewApp() *App {
	app := &App{crawler: crawler.NewLCKSetResultCrawler()}
	app.teamMatcher = matcher.NewLCKTeamMatcher()
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
	app.Repo = repository.NewFileRepository(repository.DEFAULT_RECORDING_JSON_FILE_NAME)
	app.server = api.NewHTTPServer()
	return app
}

func (app *App) MakeLCKVideoItemList() videoitem.VideoItemListMapper {
	r, _ := regexp.Compile("(0[1-9]|1[0-2]).(0[1-9]|[12][0-9]|3[01])")
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
			res := string(r.Find([]byte(videoItem.Snippet.Title)))
			if res == "" {
				fmt.Printf("debug %s find:%v\n", videoItem.Snippet.Title, res)
				continue
			}
			fmt.Println(videoItem.Snippet.Title, ": ", res)
			monthDay := strings.Split(res, ".")
			if len(monthDay) != 2 {
				fmt.Printf("debug monthDay:%v\n", monthDay)
				continue
			}
			date := fmt.Sprintf("%v-%s-%s", videoItem.Snippet.PublishedAt.Year(), monthDay[0], monthDay[1])
			// date := videoItem.Snippet.PublishedAt.Format("2006-01-02")
			videoItemMapper[date] = append(videoItemMapper[date], videoitem.NewVideoItem(
				playListItem.Snippet.Title,
				videoItem.Snippet.Title,
				videoItem.Snippet.ResourceID.VideoID,
				videoItem.Snippet.PublishedAt))
		}
	}
	return videoItemMapper
}

func main() {
	app := NewApp()
	videoItemMapper := app.MakeLCKVideoItemList()
	repoDataList := repository.LCKMatchListModel{}
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
			matchModel.PlayList = item.PlayList
			matchModel.Title = item.Title
			matchModel.LCKMatchVideoModel = *repository.NewLCKMatchVideoModel(item.PlayList, item.Title, item.VideoId, k)
			fmt.Println("playlist:", item.PlayList)
			fmt.Println("title:", item.Title)
			fmt.Println("video:", item.VideoId)
			if len(setResultData) > i {
				matchModel.LCKMathTeamModel = *repository.NewLCKMathTeamModel(
					setResultData[i].TeamScore1.Team,
					setResultData[i].TeamScore1.Score,
					setResultData[i].TeamScore2.Team,
					setResultData[i].TeamScore2.Score)
				fmt.Println("team1:", setResultData[i].TeamScore1.Team)
				fmt.Println("team1-result:", setResultData[i].TeamScore1.Score)
				fmt.Println("team2:", setResultData[i].TeamScore2.Team)
				fmt.Println("team2-result:", setResultData[i].TeamScore2.Score)
			} else {
				matchModel.IsError = true
				fmt.Println("Error ", setResultData, " ", i)
			}
			fmt.Println("date:", item.Date)
			// team1 := app.teamMatcher.Match(setResultData[i].TeamScore1.Team)
			// team2 := app.teamMatcher.Match(setResultData[i].TeamScore2.Team)
			fmt.Println("------------------------")
			repoDataList.Data = append(repoDataList.Data, matchModel)
		}
	}

	if len(repoDataList.Data) == 0 {
		repoDataList.Error = "Error: There are no data"
	}
	data, err := json.MarshalIndent(repoDataList, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = app.Repo.Store(string(repository.ALL_MATCH), string(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	app.server.StartServer()
}
