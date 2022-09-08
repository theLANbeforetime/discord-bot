package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	Token             string
	Prefix            string
	TwitchClientID    string
	TwitchAccessToken string
	config            *configStruct
)

type configStruct struct {
	Token             string `json:"discordtoken"`
	Prefix            string `json:"discordprefix"`
	TwitchClientID    string `json:"twitchclientid"`
	TwitchAccessToken string `json:"twitchaccesstoken"`
}

func ReadConfig() error {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	Token = config.Token
	Prefix = config.Prefix
	TwitchClientID = config.TwitchClientID
	TwitchAccessToken = config.TwitchAccessToken

	return nil
}
