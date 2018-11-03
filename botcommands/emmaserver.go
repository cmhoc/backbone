package botcommands

import (
	"github.com/bwmarrin/discordgo"
)

var flag int = 0
var channelid = "410180831412355072" //real channel
//var channelid = "506234472866250755" //testing channel

func Flag(discord *discordgo.Session, message *discordgo.MessageCreate){
	if message.ChannelID == channelid {
		if message.Content == (commandPrefix + "flag") {
			discord.ChannelMessageSend(message.ChannelID, "Support Turned back On")
			flag = 0
		}
	}
}

//function for emmas server to autorespond in the crisis channel
func Emmaserver(discord *discordgo.Session, message *discordgo.MessageCreate){
	//time := 0

	if message.Author.ID == discord.State.User.ID {
		return
	}

	role := "<@&506244281951059978>"

	if message.ChannelID == channelid {
		if message.Content != (commandPrefix + "flag"){
			if flag == 0 {
				discord.ChannelMessageSend(message.ChannelID,
					("Hello, you've asked for help in #crisis-support! " +
						role + " will be here to help momentarily!"))
				discord.ChannelMessageSend(message.ChannelID, "Some resources that might help you are:")
				discord.ChannelMessageSend(message.ChannelID, "https://docs.google.com/document/d/1sjanD5oATaJEFAmMyEwapq-eSd3zOpfkYjUzlwldsiw/edit?usp=sharing")
			} else {
				return
			}
		}
		flag = 1
	}

	/*
	for flag == 1 {
		time := Cooldown(2)
		if time == 0 {flag = 0}
		} */
}