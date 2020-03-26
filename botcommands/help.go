package botcommands

import (
	"backbone/tools"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

func Help(discord *discordgo.Session, message *discordgo.MessageCreate) {
	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Color:       0x696969, // Grey
		Description: "Contact verielthewolf@protonmail.com if you have issues",
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:   "./help",
				Value:  "What youre currently viewing",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "./hereboy",
				Value:  "Call the puppy to you",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "./pet",
				Value:  "Pet the puppy",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "./fetch",
				Value:  "Throw a ball for the puppy",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "./goodboy",
				Value:  "Shows a random dog gif",
				Inline: true,
			},
			&discordgo.MessageEmbedField{
				Name:   "./shakeapaw",
				Value:  "Shake a paw!",
				Inline: true,
			},
		},
		//Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title: "A List of non-Server Specific Commands",
	}

	if message.Content == "./help" {
		_, err := discord.ChannelMessageSendEmbed(message.ChannelID, embed)
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
	}
}

//somethings broken here with the for loop. It literally worked before and it wasnt changed
func Cmhochelp(discord *discordgo.Session, message *discordgo.MessageCreate) {
	var temp string
	server, _ := discord.Guild(Serverid)
	tools.Log.WithFields(logrus.Fields{"#Channels": len(server.Channels)}).Debug("CMHoC Channels")

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x696969, // Grey
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "./seats",
				Value:  "WIP",
				Inline: true,
			},
		},
		Title: "A List of CMHoC Specific Commands",
	}

	if message.Content == commandPrefix+"helpcmhoc" {
		for i := 0; i < len(server.Channels); i++ {
			temp = server.Channels[i].ID
			if message.ChannelID == temp {
				_, err := discord.ChannelMessageSendEmbed(message.ChannelID, embed)
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
			}
		}
	}
}
