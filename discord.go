package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.ChannelID == "651000959446548480" && m.Content != "!clearMOTD" {
		messageOfTheDay = m.Content
	}
	if m.Content == "!clearMOTD" {
		messageOfTheDay = ""
	}
	if m.ChannelID == "652778441833447424" && m.Content != "!clearSchoolBot" {
		schoolbot = m.Content
	}
	if m.ChannelID == "682341285977260052" && m.Content != "!clearJson" {
		jsonData = separation(m.Content)
	}
	if m.Content == "!clearSchoolBot" {
		schoolbot = ""
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
		fmt.Println("Print channel id: ", m.ChannelID)
	}

}

func separation(ctx string) string {
	complete := strings.Split(ctx, "@@@")
	jsonout, _ := json.Marshal(complete)
	return string(jsonout)
}
