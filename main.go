/*
Initially as a test for discordgo, this test has grown into the main backbone for
cmhocs automation process and controls the majority if not the entirety. Because all streams of the automation are
connected, much of it can be accessed and adjusted through the discord bot or the website.
For detailed documentation please view the readme's in each directory.
Sensitive information has been omitted from the public release.
Author: /u/thehowlinggreywolf
Contact @: verielthewolf@gmail.com
Copyright: MIT License
*/
package main

import (
	"discordbot/botcommands"
	"discordbot/logging"
	"discordbot/sql-interface"
	"discordbot/webserver"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/handlers"
	"net/http"
)

const ( // General constants for connecting to the bot
	secret   = ""
	botid    =
	bottoken = ""
	userid   = ""
)

func main() {
	//connecting to the database
	db := database.Login()
	db.Begin()
	//pulling bill data
	bills := database.Billsr(db)
	logger.Log.WithField("Bill", bills).Info("Bills Loaded")

	//Followings is all for the discord bot
	discord, err := discordgo.New("Bot " + bottoken)
	if err != nil {
		logger.Log.Panic("Could not create discord bot")
	}
	//loading the bot functions as a goroutine
	go bothandler(discord)
	//opening the connection to discord
	err = discord.Open()
	if err != nil {
		logger.Log.Panic("Error Connecting to Discord")
	}
	//stops the discord connection from closing until the main script is finished running
	defer discord.Close()

	//The following is code for the webserver
	//Open the webserver
	staticFileDirectory := http.Dir("./webserver/static/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	http.Handle("/", staticFileHandler)
	err = http.ListenAndServe(":8080", handlers.LoggingHandler(logger.Log.Out, &myHandler{}))
	//closes the program if it fails to create the server
	if err != nil {
		logger.Log.Fatal("Could not create webserver")
	}
}

func bothandler(discord *discordgo.Session) {
	//adding the commands to the discord bot
	discord.AddHandler(botcommands.Messagelog) //This just outputs messages sent to the log. Used to debug
	discord.AddHandler(botcommands.Emmaserver)
	discord.AddHandler(botcommands.Hereboy)
	discord.AddHandler(botcommands.Pet)
	discord.AddHandler(botcommands.Flag)
	discord.AddHandler(botcommands.Fetch)
	discord.AddHandler(botcommands.Help)
	discord.AddHandler(botcommands.Goodboy)
	discord.AddHandler(botcommands.Eatthepuppy)
	discord.AddHandler(botcommands.Shakeapaw)
	discord.AddHandler(botcommands.Cmhochelp)
	discord.AddHandler(botcommands.Seatchart)
	discord.AddHandler(botcommands.Todo)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err := discord.UpdateStatus(0, "Use ./help if youre stuck!")
		if err != nil {
			logger.Log.Panic("Error Setting Discord Status")
		}
		servers := discord.State.Guilds
		logger.Log.Printf("Started on %d servers", len(servers))
	})
	//returns true once run
	logger.Log.Info("Discord Functions Loaded")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, err := webserver.Handlers[r.URL.String()]; err {
		h(w, r)
		return
	}
}
