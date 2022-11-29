package main

import (
	"fmt"
	"net/http"

	discord "github.com/bwmarrin/discordgo"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type messageToSend struct {
	Message string `json:"Message"`
	Value   string `json:"Value"`
}

var messageOfTheDay string
var schoolbot string
var jsonData string

func main() {
	// Create a new Discord session using the provided bot token.
	dg, err := discord.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, messageOfTheDay)
	})

	e.GET("/Telebot", func(c echo.Context) error {
		return c.String(http.StatusOK, schoolbot)
	})

	e.POST("/callme", func(c echo.Context) error {
		u := new(messageToSend)
		err := c.Bind(u)
		if err != nil {
			fmt.Println("Bind error: ", err)
		}
		if u.Value == "call" {
			dg.ChannelMessageSend("650985405788979220", u.Message)
		}
		if u.Value == "feedback" {
			dg.ChannelMessageSend("648413092526686219", u.Message)
		}
		return c.JSON(http.StatusOK, u)
	})

	e.GET("/jsondata", func(c echo.Context) error {
		return c.String(http.StatusOK, jsonData)
	})

	// Start server
	e.Start(":3000")
}
