package botcommands

import (
	"github.com/bwmarrin/discordgo"
	"time"
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
		},
		Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
		Title:     "A List of non-Server Specific Commands",
	}

	if  message.Content == "./help" {
		discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	}
}