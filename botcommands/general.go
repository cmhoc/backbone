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
		discord.ChannelMessageSend(message.ChannelID, send)
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

		discord.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Color:  0x696969, // Grey
			Image: &discordgo.MessageEmbedImage{
				URL: "https://media.giphy.com/media/" + gif.Id + "/giphy.gif",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		})
		tools.Log.Trace("Goodboy Gif Embed")
	}
}

func Eatthepuppy(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Content == commandPrefix+"eatthepuppy" {
		user := message.Author.Mention()
		send := "?ban " + user + " No eating the puppy"
		discord.ChannelMessageSend(message.ChannelID, send)
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
		discord.ChannelMessageSendEmbed(message.ChannelID, embed)
	}
}

func Todo(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, commandPrefix+"todo") {
		if strings.Contains(message.Content, commandPrefix+"todo list") {
			return
		}
		if strings.Contains(message.Content, commandPrefix+"todo delete") {
			return
		}
		if strings.TrimPrefix(message.Content, commandPrefix+"todo") == "" {
			return
		}
		if message.Author.ID == "155084706868625408" {
			//trimming prefix to just get the message
			temp := strings.TrimPrefix(message.Content, commandPrefix+"todo ")
			tools.Log.WithField("Addition", temp).Trace("Todo Addition")
			//outputting to the files
			output, err := os.OpenFile("todo", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				tools.Log.Debug("Error Opening File")
				discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
			}
			defer output.Close()
			output.WriteString("[" + temp + "] ")
			//success message
			discord.ChannelMessageSend(message.ChannelID, "Todo content added "+"["+temp+"]")
		}
	}
}

func Todoread(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, commandPrefix+"todo list") {
		if message.Author.ID == "155084706868625408" {
			//reading in the file
			input, err := os.OpenFile("todo", os.O_RDONLY, 0666)
			if err != nil {
				discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
				tools.Log.Debug("Error Opening File")
				return
			}
			defer input.Close()
			info, _ := input.Stat()
			size := info.Size()
			data := make([]byte, size)
			input.Read(data)
			temp := string(data)
			output := strings.Split(temp, "[")
			tools.Log.WithField("# of Objects", len(output)).Trace("Todo List")
			discord.ChannelMessageSend(message.ChannelID, "Todo list contents:")
			for i := 1; i < len(output); i++ {
				discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%d: %s", i, "["+output[i]))
			}
		}
	}
}

func Tododelete(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, commandPrefix+"todo delete") {
		if message.Author.ID == "155084706868625408" {
			//reading in the file
			file, err := os.OpenFile("todo", os.O_RDWR, 0666)
			if err != nil {
				discord.ChannelMessageSend(message.ChannelID, "Error Opening File")
				tools.Log.Debug("Error Opening File")
				return
			}
			info, _ := file.Stat()
			size := info.Size()
			data := make([]byte, size)
			file.Read(data)
			temp := string(data)
			output := strings.Split(temp, "[")
			tools.Log.WithField("# of Objects", len(output)).Trace("Todo List")
			//determining the item
			index1 := strings.TrimPrefix(message.Content, commandPrefix+"todo delete ")
			index, err := strconv.Atoi(index1)
			//dealing with errors
			if err != nil {
				discord.ChannelMessageSend(message.ChannelID, "Int not given")
				tools.Log.Debug("Int Not Given")
				return
			}
			if (index-1 > len(output)) || (index-1 < 0) {
				discord.ChannelMessageSend(message.ChannelID, "Index not in range")
				return
			}
			//deleting the item
			discord.ChannelMessageSend(message.ChannelID, "Deleted: "+"["+output[index])
			output = append(output[:index], output[index+1:]...)

			file.Close()
			//overwriting the file
			file, _ = os.Create("todo")
			defer file.Close()
			for i := 1; i < len(output); i++ {
				file.WriteString("[" + output[i])
			}
		}
	}
}
