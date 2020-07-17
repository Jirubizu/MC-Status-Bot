package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	config Config
}

func main() {
	var config Config
	config.getConfig()
	bot := Bot{
		config: config,
	}
	dg, err := discordgo.New("Bot " + bot.config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}

	dg.AddHandler(bot.messageCreate)
	dg.AddHandler(bot.botConnect)

	if err = dg.Open(); err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
}
