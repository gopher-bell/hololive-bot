package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gopher-bell/hololive-bot/discord"
	"github.com/gopher-bell/hololive-bot/log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 시드 설정
	rand.Seed(time.Now().UnixNano())

	// 로그 설정
	if fn, err := log.SetupZap(); err != nil {
		panic(err)
	} else {
		defer fn()
	}

	session, err := discord.New()
	if err != nil {
		panic(err)
	}

	session.AddHandler(messageCreate)

	err = session.Open()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	fmt.Println(session.State.Guilds[0])

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.ZapLog.Errorw("interrupt occurred shutting down server")
	/*
		Do Something
	*/
	log.ZapLog.Infow("gracefully shut downed server")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println("guild", m.GuildID)
	fmt.Println("channel", m.ChannelID)
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
