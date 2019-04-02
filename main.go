/*
Initially as a test for discordgo, this test has grown into the main backbone for
cmhocs automation process and controls the majority if not the entirety. Because all streams of the automation are
connected, much of it can be accessed and adjusted through the discord bot or the website.
Sensitive information has been omitted from the public release.
Author: /u/thehowlinggreywolf
Contact @: verielthewolf@gmail.com
Copyright: MIT License
*/
package main

import (
	dcom "backbone/botcommands"
	"backbone/google-interface"
	"backbone/sql-interface"
	"backbone/tools"
	"backbone/webserver"
	"github.com/bwmarrin/discordgo"
	"net/http"
)

func main() {
	//connecting to the database
	db := database.Login()
	_, err := db.Begin()
	if err != nil {
		tools.Log.WithField("Error", err).Debug("Error Loading Database")
	}

	//Loading the Automatic Updater
	go database.DBUpdating(db)

	//Initial Updates
	err = database.ForceUpdate(db)
	if err != nil {
		tools.Log.WithField("Error", err).Debug("Could Not Update SQL based vars")
	}

	//Loading google Authentication
	err = google.GoogleAuth()
	if err != nil {
		tools.Log.Debug("Error Identified with Google Authentication, all related functions will not work.")
	} else {
		tools.Log.Info("Google API services loaded.")
	}

	//Followings is all for the discord bot
	discord, err := discordgo.New("Bot " + dcom.Bottoken)
	if err != nil {
		tools.Log.Panic("Could not create discord bot")
	}
	//loading the bot functions as a goroutine
	if tools.Conf.GetBool("discord") {
		go bothandler(discord)
	} else {
		tools.Log.Info("Discord Functions not Loaded")
	}
	//opening the connection to discord
	err = discord.Open()
	if err != nil {
		tools.Log.Panic("Error Connecting to Discord")
	}
	//stops the discord connection from closing until the main script is finished running
	defer func() {
		err := discord.Close()
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Error Closing Discord")
			return
		}
	}()

	//Turning on the clear terminal loop if debug mode is on
	if tools.Conf.GetBool("debug") {
		go tools.ClearLoop()
	}

	//The following is code for the webserver
	//Open the webserver
	err = webserving()
	//closes the program if it fails to create the server
	if err != nil {
		tools.Log.WithField("Error", err).Fatal("Error Creating Webserver")
	}
}

//adds all the handlers to the bot
func bothandler(discord *discordgo.Session) {
	discord.AddHandler(dcom.Messagelog) //This just outputs messages sent to the log. Used to debug
	discord.AddHandler(dcom.Emmaserver)
	discord.AddHandler(dcom.Hereboy)
	discord.AddHandler(dcom.Pet)
	discord.AddHandler(dcom.Flag)
	discord.AddHandler(dcom.Fetch)
	discord.AddHandler(dcom.Help)
	discord.AddHandler(dcom.Goodboy)
	discord.AddHandler(dcom.Eatthepuppy)
	discord.AddHandler(dcom.Shakeapaw)
	discord.AddHandler(dcom.Cmhochelp)
	discord.AddHandler(dcom.Seatchart)
	discord.AddHandler(dcom.Todo)
	discord.AddHandler(dcom.BillSub)
	discord.AddHandler(dcom.VoteCount)
	discord.AddHandler(dcom.Nyoom)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err := discord.UpdateStatus(0, "Visit https://cmhoc.com!")
		if err != nil {
			tools.Log.Debug("Error Setting Discord Status")
		}
		servers := discord.State.Guilds
		tools.Log.Printf("Started on %d servers", len(servers))
	})
	//returns true once run
	tools.Log.Info("Discord Functions Loaded")
}

//launches the webserver
func webserving() error {

	//creating a webserver
	mux := http.NewServeMux()

	//static file serving
	staticFileDirectory := http.Dir(tools.Conf.GetString("wfiles"))
	staticFileHandler := http.StripPrefix("/", http.FileServer(staticFileDirectory))
	mux.Handle("/", staticFileHandler)

	//handler scripts
	//api scripts
	mux.HandleFunc("/api/billdata", webserver.Billsjson)
	mux.HandleFunc("/api/partydata", webserver.Partyjson)
	mux.HandleFunc("/api/votedata", webserver.Votejson)
	mux.HandleFunc("/api/anounce", webserver.AnnouncementRSS)
	//auth scripts
	mux.HandleFunc("/auth/user", webserver.AuthSend)

	err := http.ListenAndServeTLS(tools.Conf.GetString("wdomain")+":"+tools.Conf.GetString("wport"), tools.Conf.GetString("wcert"),
		tools.Conf.GetString("wkey"), webserver.Logging(mux))
	if err != nil {
		tools.Log.Debug("Error in TLS Serving, Serving Without.")
		err = http.ListenAndServe(tools.Conf.GetString("wdomain")+":"+tools.Conf.GetString("wport"), webserver.Logging(mux))
		if err != nil {
			return err
		}
	}

	return nil
}
