package main

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

type MCResponse struct {
	MinecraftNet     string `json:"minecraft.net"`
	SessionMinecraft string `json:"session.minecraft.net"`
	AccountMojang    string `json:"account.mojang.com"`
	AuthServMojang   string `json:"authserver.mojang.com"`
	SessionMojang    string `json:"sessionserver.mojang.com"`
	APIMojang        string `json:"api.mojang.com"`
	TextureMinecraft string `json:"textures.minecraft.net"`
	Mojang           string `json:"mojang.com"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if strings.Fields(m.Content)[0] != "rules" {
		return
	}

	// Create a command system to set channels
}

func botConnect(s *discordgo.Session, c *discordgo.Connect) {
	file, err := os.Create("status.cache")

	if err != nil {
		fmt.Println(err)
		return
	}

	go loop(s, file)
}

func UpdateStatus(s *discordgo.Session, file *os.File) {
	resp, err := http.Get("https://status.mojang.com/check")
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	var resStruct MCResponse

	body = []byte(strings.ReplaceAll(strings.ReplaceAll(string(body), "{", ""), "}", ""))
	body = []byte(strings.ReplaceAll(strings.ReplaceAll(string(body), "[", "{"), "]", "}"))

	err = json.Unmarshal(body, &resStruct)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadFile("status.cache")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	var cacheStruct MCResponse

	// if the file is empty write the data for initial cache.
	if len(data) == 0 {
		_, _ = file.Write(body)
		return
	} else { // else get the cached as a struct
		err = json.Unmarshal(body, &cacheStruct)
		if err != nil {
			log.Fatal(err)
		}
	}

	respValues := reflect.ValueOf(resStruct)
	cacheValues := reflect.ValueOf(cacheStruct)

	typeOfS := respValues.Type()

	var tempFields []*discordgo.MessageEmbedField

	for i := 0; i < respValues.NumField(); i++ {
		if respValues.Field(i).Interface() == cacheValues.Field(i).Interface() {
			continue
		}
		emoji, desc := convertToData(fmt.Sprintf("%v", respValues.Field(i).Interface()))
		tempFields = append(tempFields, &discordgo.MessageEmbedField{
			Name:   strings.ReplaceAll(strings.ReplaceAll(string(typeOfS.Field(i).Tag), "\"", ""), "json:", "") + emoji,
			Value:  desc,
			Inline: false,
		})
	}

	if len(tempFields) == 0 {
		return
	} else {
		_, _ = file.Write(body)
	}

	var embed = discordgo.MessageEmbed{}
	embed.Title = "Minecraft Status Update!"
	embed.Color = 0xff0000
	embed.Fields = tempFields

	_, _ = s.ChannelMessageSendEmbed("CHANNEL ID", &embed)
}

func loop(s *discordgo.Session, file *os.File) {
	timerCh := time.Tick(time.Duration(5000) * time.Millisecond)

	for range timerCh {
		UpdateStatus(s, file)
	}
}

func convertToData(status string) (string, string) {
	switch status {
	case "green":
		return "\t\U0001F7E9", "Everything is in good working condition"
	case "yellow":
		return "\t\U0001F7E8", "Some functionality might be limited"
	case "red":
		return "\t\U0001F7E5", "There is an error with the service"
	default:
		return "\tUnknown", "\tUnknown"
	}
}
