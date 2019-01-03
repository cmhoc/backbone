package botcommands

import (
	"backbone/google-interface"
	"backbone/reddit-interface"
	"backbone/tools"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

var roles map[string]string //saving the CMHoC roles into a map

func init() {
	//loading importan roles into the map
	roles = make(map[string]string)
	roles["headmod"] = "172225589355216896"
	roles["admin"] = "475144656024240138"
	roles["mp"] = "481202246621724683"
	roles["parliament"] = "480214208634683404"
}

func checkRole(server *discordgo.Guild, author string, role string) bool {
	for i := 0; i < len(server.Members); i++ {
		if server.Members[i].User.ID == author {
			for y := 0; y < len(server.Members[i].Roles); y++ {
				if server.Members[i].Roles[y] == role {
					return true
				}
			}
			break
		}
	}
	return false
}

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

func BillSub(discord *discordgo.Session, message *discordgo.MessageCreate) {
	//loading which server is cmhoc for role checking
	//server, _ := discord.Guild(Serverid)
	var bill google.Upload
	if message.Content == commandPrefix+"submit" {
		channel, err := discord.Channel(message.ChannelID)
		if err != nil {
			tools.Log.Debug("Bill Submission Failed")
			discord.ChannelMessageSend(message.ChannelID, "Error Submitting Bill")
			return
		}
		if channel.Type == discordgo.ChannelTypeDM {

			discord.ChannelMessageSend(message.ChannelID, "Validating Role, This may take some time")
			server, _ := discord.Guild(Serverid)
			hasRole := checkRole(server, message.Author.ID, roles["mp"])

			if !hasRole {
				discord.ChannelMessageSend(message.ChannelID, "Failed to Validate role, You must be an MP to Submit Bills")
				return
			}

			if hasRole {

				discord.ChannelMessageSend(message.ChannelID, "Hello You've Indicated you want to submit a bill!")
				discord.ChannelMessageSend(message.ChannelID, "If at any time you'd like to quit, type \"exit\"")
				bill.Timestamp = time.Now().Format(time.RFC850)
				discord.ChannelMessageSend(message.ChannelID, "What is the reddit /u/ of the Bill's Author?")
				//last messages timestamp
				messages, _ := discord.ChannelMessages(channel.ID, 1, "", "", "")
				lastTime := messages[0].Timestamp

				for true { //wait until the last message is not the bots
					messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
					if messages[0].Timestamp != lastTime {
						lastTime = messages[0].Timestamp
						break
					}
				}
				if messages[0].Content == "exit" {
					discord.ChannelMessageSend(message.ChannelID, "Exiting Submission")
					return
				}
				bill.Author = messages[0].Content

				discord.ChannelMessageSend(message.ChannelID, "Please specifify if it is a \"Motion\" or \"Bill\"")
				//last messages timestamp
				messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
				lastTime = messages[0].Timestamp
				for true { //wait until the last message is not the bots
					messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
					if messages[0].Timestamp != lastTime {
						lastTime = messages[0].Timestamp
						break
					}
				}
				if messages[0].Content == "exit" {
					discord.ChannelMessageSend(message.ChannelID, "Exiting Submission")
					return
				}
				bill.Type = messages[0].Content

				discord.ChannelMessageSend(message.ChannelID, "Co-Sponsor of the bill? Use N/A if its no one")
				//last messages timestamp
				messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
				lastTime = messages[0].Timestamp
				for true { //wait until the last message is not the bots
					messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
					if messages[0].Timestamp != lastTime {
						lastTime = messages[0].Timestamp
						break
					}
				}
				if messages[0].Content == "exit" {
					discord.ChannelMessageSend(message.ChannelID, "Exiting Submission")
					return
				}
				bill.Sponsor = messages[0].Content

				discord.ChannelMessageSend(message.ChannelID, "Slot it's being submitted to? IE Government or NDP")
				//last messages timestamp
				messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
				lastTime = messages[0].Timestamp
				for true { //wait until the last message is not the bots
					messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
					if messages[0].Timestamp != lastTime {
						lastTime = messages[0].Timestamp
						break
					}
				}
				if messages[0].Content == "exit" {
					discord.ChannelMessageSend(message.ChannelID, "Exiting Submission")
					return
				}
				bill.Slot = messages[0].Content

				discord.ChannelMessageSend(message.ChannelID, "Link to the bill")
				//last messages timestamp
				messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
				lastTime = messages[0].Timestamp
				for true { //wait until the last message is not the bots
					messages, _ = discord.ChannelMessages(channel.ID, 1, "", "", "")
					if messages[0].Timestamp != lastTime {
						lastTime = messages[0].Timestamp
						break
					}
				}
				if messages[0].Content == "exit" {
					discord.ChannelMessageSend(message.ChannelID, "Exiting Submission")
					return
				}
				bill.Link = messages[0].Content

				bill.Meta = "Submitted through puppy by: " + message.Author.Username
				//submitting the completed 'Upload' struct
				err = google.GoogleBillUp(bill)
				if err != nil {
					tools.Log.Debug("Bill Submission Failed")
					discord.ChannelMessageSend(message.ChannelID, "Error Submitting Bill")
					return
				}
				discord.ChannelMessageSend(message.ChannelID, "Bill Successfully Submitted")
			}
		}
	}
}

func VoteCount(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.HasPrefix(message.Content, commandPrefix+"count") {

		discord.ChannelMessageSend(message.ChannelID, "Validating Role, This may take some time")
		server, _ := discord.Guild(Serverid)
		hasRole := checkRole(server, message.Author.ID, roles["parliament"])

		if !hasRole {
			discord.ChannelMessageSend(message.ChannelID, "Failed to Validate role, You must be an MP to Submit Bills")
			return
		}

		if hasRole {
			link := strings.Trim(message.Content, commandPrefix+"count ")

			votes, billtitles, err := reddit.Count(link)
			if err != nil {
				tools.Log.WithField("Error", err).Debug("Error Counting Votes")
				discord.ChannelMessageSend(message.ChannelID, "Error Counting Votes")
				return
			}

			err = google.GoogleVotesUp(votes, billtitles)
			if err != nil {
				tools.Log.WithField("Error", err).Debug("Error Uploading Votes")
				discord.ChannelMessageSend(message.ChannelID, "Error Uploading Votes")
				return
			}

			tools.Log.WithField("Bills", billtitles).Info("Votes Counted")

			discord.ChannelMessageSend(message.ChannelID, "Votes Counted")
		}
	}
}