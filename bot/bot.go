package bot

import (
	"discord-bot/config"
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/nicklaw5/helix/v2"
)

var BotID string
var goBot *discordgo.Session

func Run() {
	// create bot session
	goBot, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal(err)
		return
	}
	// make the bot a user
	user, err := goBot.User("@me")
	if err != nil {
		log.Fatal(err)
		return
	}
	BotID = user.ID
	goBot.AddHandler(MessageHandler)
	err = goBot.Open()
	if err != nil {
		return
	}
}

// Reads in desired user search term and then utilizes helix twitch library search top 20 users and return if they are live and playing Apex Legends
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == BotID {
	}
	// Take the ! out of the command
	if strings.HasPrefix(m.Content, "!") {
		searchTerm := strings.TrimLeft(m.Content, "!")
		client, err := helix.NewClient(&helix.Options{
			ClientID:       config.TwitchClientID,
			AppAccessToken: config.TwitchAccessToken,
		})
		if err != nil {
			panic(err)
		}
		// Search Twitch.tv for the first 20 channels similar to search name
		resp, err := client.SearchChannels(&helix.SearchChannelsParams{
			First:   20,
			Channel: searchTerm,
		})
		if err != nil {
			//handle error
		}
		// Write message to Discord to show search results
		_, _ = s.ChannelMessageSend(m.ChannelID, "Below are the users who matched for Apex Legends:\n")
		for _, respons := range resp.Data.Channels {
			if respons.GameName == "Apex Legends" {
				resultFromSearch := fmt.Sprintf("Username: %+v\n Game: %+v\n Live: %+v \n https://twitch.tv/%+v\n Link to Videos --> https://twitch.tv/%+v/videos", respons.DisplayName, respons.GameName, respons.IsLive, respons.DisplayName, respons.DisplayName)
				_, _ = s.ChannelMessageSend(m.ChannelID, resultFromSearch)
				_, _ = s.ChannelMessageSend(m.ChannelID, "####################################")
			}
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, "End of Search")
	}
}
