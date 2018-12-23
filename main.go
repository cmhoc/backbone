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
	"backbone/botcommands"
	"backbone/google-interface"
	"backbone/reddit-interface"
	"backbone/sql-interface"
	"backbone/tools"
	"backbone/webserver"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/handlers"
	"net/http"
)

func main() {
	//connecting to the database
	db := database.Login()
	db.Begin()
	//pulling bill data
	bills := database.Billsr(db)
	tools.Log.WithField("Bill", bills).Info("Bills Loaded")

	//Loading google Authentication
	err := google.GoogleAuth()
	if err != nil {
		tools.Log.Debug("Error Identified with Google Authentication, all related functions will not work.")
	} else {
		tools.Log.Info("Google API services loaded.")
	}

	//Loading Reddit Auth
	err = reddit.Auth()
	if err != nil {
		tools.Log.Debug("Error Identified with Reddit Authentication, all related functions will not work.")
	} else {
		tools.Log.Info("Reddit API Services loaded")
	}

	reddit.Count("https://www.reddit.com/r/cmhocvote/comments/9w65ny/1st_parl_2nd_session_house_vote_m7a1_m10/")

	//Followings is all for the discord bot
	discord, err := discordgo.New("Bot " + botcommands.Bottoken)
	if err != nil {
		tools.Log.Panic("Could not create discord bot")
	}
	//loading the bot functions as a goroutine
	go bothandler(discord)
	//opening the connection to discord
	err = discord.Open()
	if err != nil {
		tools.Log.Panic("Error Connecting to Discord")
	}
	//stops the discord connection from closing until the main script is finished running
	defer discord.Close()

	//Turning on the clear terminal loop if debug mode is on
	if tools.Conf.GetBool("debug") {
		go tools.ClearLoop()
	}

	//The following is code for the webserver
	//Open the webserver
	staticFileDirectory := http.Dir("./webserver/static/")
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	http.Handle("/", staticFileHandler)
	err = http.ListenAndServe(":"+tools.Conf.GetString("wport"), handlers.LoggingHandler(tools.Log.Out, &myHandler{}))
	//closes the program if it fails to create the server
	if err != nil {
		tools.Log.Fatal("Could not create webserver")
	}
}

func bothandler(discord *discordgo.Session) {
	//adding the commands to the discord bot
	discord.AddHandler(botcommands.Messagelog) //This just outputs messages sent to the log. Used to debug
	//discord.AddHandler(botcommands.Emmaserver)
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
	discord.AddHandler(botcommands.Todoread)
	discord.AddHandler(botcommands.Tododelete)
	discord.AddHandler(botcommands.BillSub)
	discord.AddHandler(botcommands.VoteCount)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err := discord.UpdateStatus(0, "Use ./help if youre stuck!")
		if err != nil {
			tools.Log.Debug("Error Setting Discord Status")
		}
		servers := discord.State.Guilds
		tools.Log.Printf("Started on %d servers", len(servers))
	})
	//returns true once run
	tools.Log.Info("Discord Functions Loaded")
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, err := webserver.Handlers[r.URL.String()]; err {
		h(w, r)
		return
	}
}
