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
	Token string="TOKEN"
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

	if m.Content == "時間" {
		t := time.Now()
		t2 := time.Date(2001, 2, 1, 3, 34, 0 , 0,  time.UTC)
		text := fmt.Sprintf("%d時%d分%d秒", t.Hour(), t.Minute(), t.Second())
		diff := t2.Sub(t)
		s.ChannelMessageSend(m.ChannelID, diff)
	}
}
