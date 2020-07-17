package main

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Token  string `yaml:"token"`
	ChannelId string `yaml:"channelID"`
	ServerIp string `yaml:"serverIP"`
}

func (config *Config) getConfig() *Config{
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return config
}
