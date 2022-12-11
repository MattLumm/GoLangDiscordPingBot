package main

import(
	"fmt"
	"github.com/MattLumm/GoLangDiscordPingBot/bot"
	"github.com/MattLumm/GoLangDiscordPingBot/config" //NOTE: Token Removed for protection on the files uploaded to GitHub
)

func main(){
	err := config.ReadConfig()

	if err !=nil{
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}