package main

import (
	"fmt"

	"github.com/gyu-young-park/lck_data_generator/channel"
	"github.com/gyu-young-park/lck_data_generator/config"
	"github.com/gyu-young-park/lck_data_generator/playlist"
	playlistitems "github.com/gyu-young-park/lck_data_generator/playlistItems"
)

type App struct{
	Config *config.Config
	ChannelService channel.Service
	PlayListService playlist.Service
	PlayListItemsService playlistitems.Service
}

func NewApp() *App {
	app := &App{}
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
	channelId, err := app.ChannelService.GetChannelId()
	if err != nil {
		panic(err)
	}
	app.PlayListService = playlist.NewServiceWithChannelId(app.Config.Key,channelId)
	playlistIds, err := app.PlayListService.GetPlayListId()
	if err != nil {
		panic(err)
	}
	app.PlayListItemsService = playlistitems.NewServiceWithPlayListId(app.Config.Key)
	for _, id := range playlistIds {
		res, err := app.PlayListItemsService.GetPlayListItems(id)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	}
	return app
}

func main(){
	app := NewApp()
	fmt.Println(app)
}