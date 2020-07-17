package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func (bot *Bot) messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Fields(m.Content)[0] != "rules" {
		return
	}

	// Create a command system to set channels
}

func (bot *Bot) botConnect(s *discordgo.Session, c *discordgo.Connect) {
	file, err := os.Create("status.cache")

	if err != nil {
		fmt.Println(err)
		return
	}

	go bot.loop(s, file)
}

func (bot *Bot) UpdateStatus(s *discordgo.Session, file *os.File) {
	mojangStatusBody := getResponse("https://status.mojang.com/check")

	mcServerBody := getResponse(fmt.Sprintf("https://api.mcsrvstat.us/2/%s", bot.config.ServerIp))

	var mcServerStruct MCSrvStatus

	err := json.Unmarshal(mcServerBody, &mcServerStruct)
	if err != nil {
		mcServerStruct = MCSrvStatus{
			Ip:       "",
			Port:     0,
			Debug:    DebugStruct{},
			MOTD:     MOTDStruct{},
			Players:  PlayersStruct{},
			Version:  "",
			Online:   false,
			Protocal: 0,
			Hostname: "",
			Icon:     "",
			Software: "",
		}
	}

	var resStruct ResSrvStatus

	mojangStatusBody = []byte(strings.ReplaceAll(strings.ReplaceAll(string(mojangStatusBody), "{", ""), "}", ""))
	mojangStatusBody = []byte(strings.ReplaceAll(strings.ReplaceAll(string(mojangStatusBody), "[", "{"), "]", fmt.Sprintf(",\"online\":\"%t\"}", mcServerStruct.Online)))

	err = json.Unmarshal(mojangStatusBody, &resStruct)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("status.cache")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var cacheStruct ResSrvStatus

	// if the file is empty write the data for initial cache.
	if len(data) == 0 {
		_, _ = file.Write(mojangStatusBody)
		return
	} else { // else get the cached as a struct
		err = json.Unmarshal(mojangStatusBody, &cacheStruct)
		if err != nil {
			log.Fatal(err)
		}
	}

	tempFields := embedGenerator(resStruct, cacheStruct)

	if len(tempFields) == 0 {
		return
	} else {
		_, _ = file.Write(mojangStatusBody)
	}

	var embed = discordgo.MessageEmbed{}
	embed.Title = "Minecraft Status Update!"
	embed.Color = 0xff0000
	embed.Fields = tempFields

	_, _ = s.ChannelMessageSendEmbed(bot.config.ChannelId, &embed)
}

func (bot *Bot) loop(s *discordgo.Session, file *os.File) {
	timerCh := time.Tick(time.Duration(5000) * time.Millisecond)

	for range timerCh {
		bot.UpdateStatus(s, file)
	}
}
