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
	"discordbot/sql-interface"
	"discordbot/logging"
	"discordbot/webserver"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const ( // General constants for connecting to the bot

)

//initialization function
//having issues running this on the EC2 server. Need to look into that
func init() {
	// tries to login to the database. If it fails the program exits.
	if (database.Login()) == true {
		logger.Log.Info("Connected to CMHoC's Database")
	} else {
		logger.Log.Fatal("Could not connect to the Database")
	}
}

func main() {
	//starting up the web server
	web := mux.NewRouter()

	//Followings is all for the discord bot
	discord, err := discordgo.New("Bot " + bottoken)
	if err != nil {logger.Log.Panic("Could not create discord bot")}
	//loading the bot functions
	if bothandler(discord) {logger.Log.Info("Discord Functions Loaded")}
	//opening the connection to discord
	err = discord.Open()
	if err != nil {logger.Log.Panic("Error Connecting to Discord")}
	//stops the discord connection from closing until the main script is finished running
	defer discord.Close()

	//The following is code for the webserver
	//webserver serving files
	web.Handle("/", http.FileServer(http.Dir("./webserver/")))
	web.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webserver/"))))

	//webserver handlers
	web.Handle("/vote", webserver.BillHandler).Methods("GET")
	web.Handle("/vote/{bill}", webserver.AddVoteHandler).Methods("POST")

	//Open the webserver
	err = http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, web))
	//closes the program if it fails to create the server
	if err != nil {logger.Log.Fatal("Could not create webserver")}

	//makes the main script run indefinitely
	<-make(chan struct{})

}

func bothandler(discord *discordgo.Session) bool {
	//adding the commands to the discord bot
	discord.AddHandler(botcommands.Messagelog) //This just outputs messages sent into the console. Used primarily to debug
	discord.AddHandler(botcommands.Emmaserver)
	discord.AddHandler(botcommands.Hereboy)
	discord.AddHandler(botcommands.Pet)
	discord.AddHandler(botcommands.Flag)
	discord.AddHandler(botcommands.Fetch)
	discord.AddHandler(botcommands.Help)
	discord.AddHandler(botcommands.Goodboy)
	discord.AddHandler(botcommands.Eatthepuppy)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err := discord.UpdateStatus(0, "Use ./help if youre stuck!")
		if err != nil {logger.Log.Panic("Error Setting Discord Status")}
		servers := discord.State.Guilds
		fmt.Printf("Started on %d servers\n", len(servers))
	})
	//returns true once run
	return true
}