package botcommands

import (
	"backbone/tools"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/paddycarey/gophy"
	"github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//constants used for ALL of the discord package (Vars also all of the discord package)
const (
	commandPrefix = "./"
	Serverid      = "172223256277942273"
)

var ( // General constants for connecting to the bot
	secret   = tools.Conf.GetString("secret")
	botid    = tools.Conf.GetInt("botid")
	Bottoken = tools.Conf.GetString("bottoken")
	userid   = tools.Conf.GetString("userid")
)

var server *discordgo.Guild

//simple testing command
func Hereboy(discord *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == discord.State.User.ID {
		return
	}

	send := "Woof!"

	if message.Content == (commandPrefix + "hereboy") {
		_, err := discord.ChannelMessageSend(message.ChannelID, send)
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
	}
}

//messagelog
func Messagelog(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	server := discord.State.Guilds[0]
	if user.ID == userid || user.Bot {
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
	tools.Log.WithFields(logrus.Fields{
		"Message":  message.Content,
		"Sever":    server.Name,
		"ServerID": server.ID,
		//"Author": message.Author,
		//"Channel": message.ChannelID,
		//"Time": message.Timestamp,
	}).Debug("Message Log")
	//fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author) // old tools
}

func Pet(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	send := "^.^ *wags tail*"

	if message.Content == (commandPrefix + "pet") {
		_, err := discord.ChannelMessageSend(message.ChannelID, send)
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
	}
}

func Fetch(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == (commandPrefix + "fetch") {
		rand.Seed(time.Now().UTC().UnixNano())
		temp := rand.Intn(2)
		if temp == 0 {
			_, err := discord.ChannelMessageSend(message.ChannelID, "*wags tail* Want to throw again?")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
		} else if temp == 1 {
			_, err := discord.ChannelMessageSend(message.ChannelID, "*whines* It looks like the puppy couldnt find the ball!")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
		}
	}
}

func Goodboy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	//using gophy to get dog gifs
	if message.Content == commandPrefix+"goodboy" {
		rand.Seed(time.Now().UTC().UnixNano())
		limit := 50
		co := &gophy.ClientOptions{}
		client := gophy.NewClient(co)
		gif, err := client.GetGifById("ygCJ5Bul73NArGOSFN")
		if err != nil {
			tools.Log.Error("Problem Initializing Gifs")
		}
		gifs, numgifs, err := client.SearchGifs("dog", "pg", limit, 0)
		if err != nil {
			tools.Log.Error("Problem Loading Gifs")
		} else {
			tools.Log.WithFields(logrus.Fields{
				"Gifs":               len(gifs),
				"#Reportedly Loaded": numgifs,
			}).Debug("Number of Gifs loaded")
		}
		count := rand.Intn(len(gifs))
		gif = gifs[count]
		tools.Log.WithField("gif#", count).Trace("Gif Set for embed")

		_, err = discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x696969, // Grey
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media.giphy.com/media/" + gif.Id + "/giphy.gif",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
		tools.Log.Trace("Goodboy Gif Embed")
	}
}

func Eatthepuppy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == commandPrefix+"eatthepuppy" {
		user := message.Author.Mention()
		send := "?ban " + user + " No eating the puppy"
		_, err := discord.ChannelMessageSend(message.ChannelID, send)
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
	}
}

func Shakeapaw(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == commandPrefix+"shakeapaw" {
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x696969, // Grey
			Image: &discordgo.MessageEmbedImage{
				URL: "https://static1.squarespace.com/static/5907d42520099e374ad11ba1/t/59197af9c534a5e1ade4a2b8/1494842118347/dog-shake-hands.jpg?format=500w",
			},
		}
		_, err := discord.ChannelMessageSendEmbed(message.ChannelID, embed)
		if err != nil {
			tools.Log.WithField("Error", err).Warn("Unusual Error")
			return
		}
	}
}

func Todo(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == "155084706868625408" {
	if strings.Contains(message.Content, commandPrefix+"todo") {
		if strings.Contains(message.Content, commandPrefix+"todo list") {
			//reading in the file
			file, err := os.OpenFile("todo", os.O_RDWR, 0666)
			if err != nil {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
				tools.Log.Debug("Error Opening File")
				return
			}
			info, _ := file.Stat()
			size := info.Size()
			data := make([]byte, size)
			_, err = file.Read(data)
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Error Reading File")
				return
			}
			temp := string(data)
			output := strings.Split(temp, "[")
			tools.Log.WithField("# of Objects", len(output)).Trace("Todo List")
			//determining the item
			index1 := strings.TrimPrefix(message.Content, commandPrefix+"todo delete ")
			index, err := strconv.Atoi(index1)
			//dealing with errors
			if err != nil {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Int not given")
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
				tools.Log.Debug("Int Not Given")
				return
			}
			if (index-1 > len(output)) || (index-1 < 0) {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Index not in range")
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
				return
			}
			//deleting the item
			_, err = discord.ChannelMessageSend(message.ChannelID, "Deleted: "+"["+output[index])
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
			output = append(output[:index], output[index+1:]...)

			err = file.Close()
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Error Closing File")
				return
			}
			//overwriting the file
			file, _ = os.Create("todo")
			defer func() {
				err := file.Close()
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Error Closing File")
					return
				}
			}()
			for i := 1; i < len(output); i++ {
				_, err := file.WriteString("[" + output[i])
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Error Writing to File")
					return
				}
			}
		}
		if strings.Contains(message.Content, commandPrefix+"todo delete") {
			//reading in the file
			input, err := os.OpenFile("todo", os.O_RDONLY, 0666)
			if err != nil {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
				tools.Log.Debug("Error Opening File")
				return
			}
			defer func() {
				err := input.Close()
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Error Closing File")
					return
				}
			}()
			info, _ := input.Stat()
			size := info.Size()
			data := make([]byte, size)
			_, err = input.Read(data)
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Error Reading Data")
				return
			}
			temp := string(data)
			output := strings.Split(temp, "[")
			tools.Log.WithField("# of Objects", len(output)).Trace("Todo List")
			_, err = discord.ChannelMessageSend(message.ChannelID, "Todo list contents:")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
			for i := 1; i < len(output); i++ {
				_, err := discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%d: %s", i, "["+output[i]))
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
			}
		}
		if strings.TrimPrefix(message.Content, commandPrefix+"todo") == "" {
			_, err := discord.ChannelMessageSend(message.ChannelID, "Invalid Usage.")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
		}
			//trimming prefix to just get the message
			temp := strings.TrimPrefix(message.Content, commandPrefix+"todo ")
			tools.Log.WithField("Addition", temp).Trace("Todo Addition")
			//outputting to the files
			output, err := os.OpenFile("todo", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				tools.Log.Debug("Error Opening File")
				_, err := discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Unusual Error")
					return
				}
			}
			defer func() {
				err := output.Close()
				if err != nil {
					tools.Log.WithField("Error", err).Warn("Error Closing File")
				}
			}()
			_, err = output.WriteString("[" + temp + "] ")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Error Writing to File")
				return
			}
			//success message
			_, err = discord.ChannelMessageSend(message.ChannelID, "Todo content added "+"["+temp+"]")
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
		}
	}
}

//A meme, reacts with 'nyoom' to messages containing 'nyoom'
func Nyoom(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(strings.ToLower(message.Content), "nyoom") {
		temp, err := discord.ChannelMessage("552232502782066729","552272535350280222") //couldnt figure out the ID's so ill have the program do it
		if err != nil {
			if err != nil {
				tools.Log.WithField("Error", err).Error("Original Nyoom Gone")
				return
			}
		}
		var reactions [5]*discordgo.Emoji
		for i:=0;i<5;i++ {
			reactions[i] = temp.Reactions[i].Emoji
		}
		var apinames [5]string
		for i:=0;i<5;i++ {
			apinames[i] = reactions[i].APIName()
		}
		for i:=0;i<5;i++ {
			err = discord.MessageReactionAdd(message.ChannelID, message.ID, apinames[i])
			if err != nil {
				tools.Log.WithField("Error", err).Warn("Unusual Error")
				return
			}
		}
	}
}