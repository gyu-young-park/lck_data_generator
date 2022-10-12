package main

import (
	"fmt"
	"gyu/channel"
	"gyu/config"
)

type App struct{
	Config *config.Config
	ChannelService channel.Service
}

func NewApp() *App {
	app := &App{}
	app.Config = config.NewConfig(config.NewConfigSetterJSON())
	app.ChannelService = channel.NewServiceWithVideoId(app.Config.Key)
	return app
}

func main(){
	app := NewApp()
	channelId, err := app.ChannelService.GetChannelId()
	if err != nil {
		panic(err)
	}
	fmt.Println(channelId)
}