package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func convertToData(status string) (string, string) {
	switch status {
	case "green":
		return "\t\U0001F7E9", "Everything is in good working condition"
	case "yellow":
		return "\t\U0001F7E8", "Some functionality might be limited"
	case "red":
		return "\t\U0001F7E5", "There is an error with the service"
	case "true":
		return "\t\U0001F7E9", "The server is back up and online"
	case "false":
		return "\t\U0001F7E5", "The server has closed down"
	default:
		return "\tUnknown", "\tUnknown"
	}
}

func embedGenerator(resStruct ResSrvStatus, cacheStruct ResSrvStatus) []*discordgo.MessageEmbedField{
	var tempFields []*discordgo.MessageEmbedField

	respValues := reflect.ValueOf(resStruct)
	cacheValues := reflect.ValueOf(cacheStruct)

	typeOfS := respValues.Type()

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
	return tempFields
}

func getResponse(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}

	return body
}