package main

import(
	"time"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var(
	Token string="Token"
)

func main(){
	dg, err := discordgo.New("Bot " + Token)
	if err != nil{
		fmt.Println("err opening connection....", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error create", err)
		return
	}

	fmt.Println("Bot is now running. Press ctrl-c to exit!!!")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	tc := time.NewTicker(time.Second * 50)

	loopContinue := true
	for loopContinue {
		select {
			case <-sc:
				loopContinue = false
				break
			case <-tc.C:
		}
	}

	dg.Close()
}


func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

	if m.Author.ID == s.State.User.ID{
		return
	}

	if m.Content == "水素" {
		text := fmt.Sprintf("あぁ～ 水素の音ォ～!!<@!%s>", m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
	}

	if m.Content == "カス" {
		text := fmt.Sprintf("お前がカス<@!%s>", m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
	}

	if m.Content == "時間" {
		t := time.Now()
		text := fmt.Sprintf("%d時%d分%d秒", t.Hour(), t.Minute(), t.Second())
		s.ChannelMessageSend(m.ChannelID, text)
	}

	if m.Content == "334" {
		JST, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println("あなたは今虚空にいます...", err)
			return
		}

		t := time.Now()
		target := time.Date(t.Year(), t.Month(), t.Day(), 3, 34, 0, 0 , JST)
		time := t.Sub(target)
		sec := 60 - int(time.Seconds())%60
		min := 60 - (int(time.Seconds()) % 3600)/60
		hour := 23 - int(time.Seconds()) / 3600
		text := fmt.Sprintf("334まで後%d時間%d分%d秒\n <@!%s>",hour, min, sec, m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
	}
}
