package botcommands

import (
	"discordbot/logging"
	"github.com/bwmarrin/discordgo"
	"github.com/paddycarey/gophy"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

var(
	userid = "402846073045123082"
	commandPrefix = "./"
)

//simple testing command
func Hereboy(discord *discordgo.Session, message *discordgo.MessageCreate){

	if message.Author.ID == discord.State.User.ID {
		return
	}

	send := "Woof!"

	if message.Content == (commandPrefix + "hereboy") {
		discord.ChannelMessageSend(message.ChannelID, send)
	}
}

//messagelog
func Messagelog(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	server := discord.State.Guilds[1]
	if user.ID == userid|| user.Bot {
		//Do nothing because the bot is talking
		return
	}
	for i := 0; i < len(discord.State.Guilds); i++ {
		for j := 0; j < len(discord.State.Guilds[i].Channels); j++ {
			if message.ChannelID == discord.State.Guilds[i].Channels[j].ID {
				server = discord.State.Guilds[i]
			}
		}
	}

	//pretty log
	logger.Log.WithFields(logrus.Fields{
		"Message": message.Content,
		"Sever": server.Name,
		//"Author": message.Author,
		//"Channel": message.ChannelID,
		//"Time": message.Timestamp,
	}).Debug("Message Log")
	//fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author) // old logging
}

func Pet(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	send := "^.^ *wags tail*"

	if message.Content == (commandPrefix + "pet") {
		discord.ChannelMessageSend(message.ChannelID, send)
	}
}

func Fetch(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == (commandPrefix + "fetch") {
		rand.Seed(time.Now().UTC().UnixNano())
		temp := rand.Intn(2)
		if temp == 0 {
			discord.ChannelMessageSend(message.ChannelID, "*wags tail* Want to throw again?")
		} else if temp == 1 {
			discord.ChannelMessageSend(message.ChannelID, "*whines* It looks like the puppy couldnt find the ball!")
		}
	}
}

func Goodboy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	//using gophy to get dog gifs
	if  message.Content == commandPrefix + "goodboy" {
		rand.Seed(time.Now().UTC().UnixNano())
		limit := 50
		co := &gophy.ClientOptions{}
		client := gophy.NewClient(co)
		gif, err := client.GetGifById("ygCJ5Bul73NArGOSFN")
		if err != nil {logger.Log.Error("Problem Initializing Gifs")}
		gifs, numgifs, err := client.SearchGifs("dog","pg",limit,0)
		if err != nil {
			logger.Log.Error("Problem Loading Gifs")
		} else {
			logger.Log.WithFields(logrus.Fields{
				"Gifs": len(gifs),
				"#Reportedly Loaded": numgifs,
			}).Debug("Number of Gifs loaded")
		}
		count := rand.Intn(len(gifs))
		gif = gifs[count]
		logger.Log.WithField("gif#", count).Trace("Gif Set for embed")

		discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Author:      &discordgo.MessageEmbedAuthor{},
			Color:       0x696969, // Grey
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media.giphy.com/media/" + gif.Id + "/giphy.gif",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		})
		logger.Log.Trace("Goodboy Gif Embed")
	}
}

func Eatthepuppy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if  message.Content == commandPrefix + "eatthepuppy" {
		user := message.Author.Mention()
		send := "?ban " + user + " No eating the puppy"
		discord.ChannelMessageSend(message.ChannelID, send)
	}
}