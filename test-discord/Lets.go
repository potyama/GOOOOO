package main

import(
	"bufio"
	"time"
	"fmt"
	"os"
	"math/rand"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)


func main(){
	rand.Seed(time.Now().UnixNano())
	Token := loadTokenFromEnv()
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

func Abs(x int) int{
	if x < 0{
		return -x
	}
	return x
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

	if m.Author.ID == s.State.User.ID{
		return
	}
	if m.Content == "!help" {
		text := fmt.Sprintf("!水素->水素の音を発します。\n!カス->罵られます。\n!time->現在時刻を表示します。\n!334->334までの時刻を表示します。\nほめて->ほめてくれるよ ")
		s.ChannelMessageSend(m.ChannelID, text)
		return
	}


	if m.Content == "!水素" {
		text := fmt.Sprintf("あぁ～ 水素の音ォ～!!<@!%s>", m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
		return
	}

	if m.Content == "!カス" {
		text := fmt.Sprintf("お前がカス<@!%s>", m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
		return
	}

	if m.Content == "!time" {
		t := time.Now()
		text := fmt.Sprintf("%d時%d分%d秒", t.Hour(), t.Minute(), t.Second())
		s.ChannelMessageSend(m.ChannelID, text)
		return
	}

	if m.Content == "!334" {
		JST, err := time.LoadLocation("Asia/Tokyo")
		if err != nil {
			fmt.Println("あなたは今虚空にいます...", err)
			return
		}

		t := time.Now()
		var target time.Time
		if(t.Hour() >=3 && t.Minute() >= 34 && t.Second() >= 1 ){
			target = time.Date(t.Year(), t.Month(), t.Day()+1, 3, 34, 0, 0 , JST)
		}else{
			target = time.Date(t.Year(), t.Month(), t.Day(), 3, 34, 0, 0 , JST)
		}
		time := t.Sub(target)

		sec := Abs(int(time.Seconds())%60)
		min := Abs((int(time.Seconds()) % 3600)/60)
		hour := Abs(int(time.Seconds()) / 3600)

		text := fmt.Sprintf("334まで後%d時間%d分%d秒\n <@!%s>",hour, min, sec, m.Author.ID)
		s.ChannelMessageSend(m.ChannelID, text)
		return
	}
	if strings.Contains(m.Content, "ほめて") == true{
		n:= rand.Intn(7)
		if n == 5{
			text := fmt.Sprintf("おう.....<@!%s>", m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, text)
			return
		}

		switch n % 3 {
		case 0:
			text := fmt.Sprintf("すごい！！！！<@!%s>", m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, text)
		case 1:
			text := fmt.Sprintf("えらい！！！！！！<@!%s>", m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, text)
		case 2:
			text := fmt.Sprintf("天才！！！！！<@!%s>", m.Author.ID)
			s.ChannelMessageSend(m.ChannelID, text)
		}
		return

	}
}

// isso love
func loadTokenFromEnv() string{
	fp, err := os.Open(".env")
	if err != nil {
		panic(err)
	}

	defer fp.Close()

	scan := bufio.NewScanner(fp)
	var Token string
	for scan.Scan(){
		Token = scan.Text()
	}
	return Token
}
