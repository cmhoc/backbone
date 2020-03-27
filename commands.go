package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/paddycarey/gophy"
	"log"
	"math/rand"
	"time"
)

const prefix string = "./"

func fetch(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if message.Author.ID == discord.State.User.ID {
		return
	}
	if message.Content == prefix + "fetch" {
		rand.Seed(time.Now().UTC().UnixNano())
		temp := rand.Intn(2)
		if temp == 0 {
			_, err := discord.ChannelMessageSend(message.ChannelID, "*wags tail* Want to throw again?")
			if err != nil {
				log.Fatal(err)
				return
			}
		} else if temp == 1 {
			_, err := discord.ChannelMessageSend(message.ChannelID, "*whines* It looks like the puppy couldnt find the ball!")
			if err != nil {
				log.Fatal(err)
				return
			}
		}
	}
}

func petPuppy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if message.Author.ID == discord.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if message.Content == prefix+"pet" {
		_, _ = discord.ChannelMessageSend(message.ChannelID, "^.^ *wags tail*")
	}

}

func goodboy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if message.Author.ID == discord.State.User.ID {
		return
	}

	//using gophy to get dog gifs
	if message.Content == prefix+"goodboy" {
		rand.Seed(time.Now().UTC().UnixNano())
		limit := 50
		co := &gophy.ClientOptions{}
		client := gophy.NewClient(co)
		//gif, err := client.GetGifById("ygCJ5Bul73NArGOSFN")
		//if err != nil {
		//log.Fatal(err)
		//}
		gifs, _, err := client.SearchGifs("dog", "pg", limit, 0)
		if err != nil {
			log.Fatal(err)
		} else {
			//TODO: log gif loading
		}
		count := rand.Intn(len(gifs))
		gif := gifs[count]

		_, err = discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x696969, // Grey
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media.giphy.com/media/" + gif.Id + "/giphy.gif",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
