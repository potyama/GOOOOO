package main

import(
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var(
	Token string="TOKEN"
)

func main(){
	dg, err := discordgo.New("Bot " + Token)
	if err != nil{
		fmt.Println("err opening connection....", err)
		return
	}

	fmt.Println("Bot is now running. Press ctrl-c to exit!!!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}


func messageCreate(s *discordgo.VoiceConnection, m *discordgo.MessageCreate){

	if m.Author.ID == s.State.user.ID{
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong")
	}

	if m.Content == "pong"{
		s.ChannelMessageSend(m.ChannelID, "Ping")
	}
}
