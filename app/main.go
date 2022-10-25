package main

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/channel"
	"github.com/gyu-young-park/lck_data_generator/config"
	"github.com/gyu-young-park/lck_data_generator/crawler"
	"github.com/gyu-young-park/lck_data_generator/matcher"
	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
	videoitem "github.com/gyu-young-park/lck_data_generator/videoItem"
)

type App struct {
	Config               *config.Config
	teamMatcher 		matcher.Matcher 
	crawler              crawler.Crawler
	ChannelService       channel.Service
	PlayListService      playlist.Service
	PlayListItemsService playlistitems.Service
}

func NewApp() *App {
	app := &App{crawler: crawler.NewLCKSetResultCrawler()}
	app.teamMatcher = matcher.NewLCKTeamMatcher()
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
	
	return app
}

// VideoItemListMapper[ date ] = item
func (app*App) MakeLCKVideoItemList() videoitem.VideoItemListMapper {
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
			date := videoItem.Snippet.PublishedAt.Format("2006-01-02")
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
	for k, v := range videoItemMapper {
		app.crawler.SetData(k)
		rawSetResultData := app.crawler.GetResult()
		setResultData, _:= rawSetResultData.([]*crawler.LCKSetDataModel)
		for i, item := range v {
			fmt.Println("------------------------")
			fmt.Println("playlist:",item.PlayList)
			fmt.Println("title:",item.Title)
			fmt.Println("video:",item.VideoId)
			if len(setResultData) > i {
				fmt.Println("team1:", setResultData[i].TeamScore1.Team)
				fmt.Println("team1-result:", setResultData[i].TeamScore1.Score)
				fmt.Println("team2:", setResultData[i].TeamScore2.Team)
				fmt.Println("team2-result:", setResultData[i].TeamScore2.Score)
			} else {
				fmt.Println("Error")
			}
			fmt.Println("date:",item.Date)
			// team1 := app.teamMatcher.Match(setResultData[i].TeamScore1.Team)
			// team2 := app.teamMatcher.Match(setResultData[i].TeamScore2.Team)
			fmt.Println("------------------------")
		}
	}
}
