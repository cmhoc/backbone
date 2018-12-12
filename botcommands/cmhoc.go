package botcommands

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func Seatchart(discord *discordgo.Session, message *discordgo.MessageCreate) {
	var temp string
	//loading which server is cmhoc for channel checking
	server, _ := discord.Guild(Serverid)

	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Color:  0x696969, // Grey
		Image: &discordgo.MessageEmbedImage{
			URL: "https://svgshare.com/i/9Jx.svg",
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if message.Content == commandPrefix+"seats" {
		for i := 0; i < len(server.Channels); i++ {
			temp = server.Channels[i].ID
			if message.ChannelID == temp {
				discord.ChannelMessageSendEmbed(message.ChannelID, embed)
			}
		}
	}
}
