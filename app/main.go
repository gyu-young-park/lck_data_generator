package main

import (
	"fmt"
	"strings"

	"github.com/gyu-young-park/lck_data_generator/channel"
	"github.com/gyu-young-park/lck_data_generator/config"
	"github.com/gyu-young-park/lck_data_generator/crawler"
	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
)

type App struct {
	Config               *config.Config
	crawler              crawler.Crawler
	ChannelService       channel.Service
	PlayListService      playlist.Service
	PlayListItemsService playlistitems.Service
}

func NewApp() *App {
	app := &App{crawler: crawler.NewLCKResultCrawler()}
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
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
	fmt.Println(playListItems)
	var tempMapper map[string]string = make(map[string]string)
	for _, playListItem := range playListItems {
		videoItems, err := app.PlayListItemsService.GetVideoItems(playListItem.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("-----------------------", playListItem.Snippet.Title, "-----------------------")
		for _, videoItem := range videoItems {
			fmt.Println("VIDEO ID:", videoItem.Snippet.ResourceID.VideoID)
			fmt.Println("PUBLISHED AT:", videoItem.Snippet.PublishedAt)
			fmt.Println("TITLE:", videoItem.Snippet.Title)
			splitedTitile := strings.Split(videoItem.Snippet.Title, "|")
			last := len(splitedTitile) - 1
			tempMapper[splitedTitile[last]] = splitedTitile[last]
			fmt.Println()
		}
		fmt.Println("-----------------------", playListItem.Snippet.Title, "-----------------------")
	}
	return app
}

func main() {
	app := NewApp()
	app.crawler.GetResult()
}
