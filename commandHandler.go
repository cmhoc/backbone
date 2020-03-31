package main

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

//While a bit redundent, used primarily to auto gen the help function
type command struct {
	function    func(s *discordgo.Session, m *discordgo.MessageCreate)
	cmd         string
	description string
}

type commandHandler struct {
	dg   *discordgo.Session
	cmds []command
}

func newCommandHandler(dg *discordgo.Session) commandHandler {
	var temp commandHandler
	temp.dg = dg
	return temp
}

func newCommand(function func(s *discordgo.Session, m *discordgo.MessageCreate), cmd string, description string) command {
	var temp command
	temp.function = function
	temp.cmd = cmd
	temp.description = description
	return temp
}

func (ch commandHandler) loadCommands() {
	//Add commands here
	ch.cmds = append(ch.cmds, newCommand(petPuppy, "pet", "Pet the puppy!"))
	ch.cmds = append(ch.cmds, newCommand(fetch, "fetch", "Play fetch with the puppy!"))
	ch.cmds = append(ch.cmds, newCommand(goodboy, "goodboy", "Return a random doggo gif"))

	//Commands just for me
	ch.cmds = append(ch.cmds, newCommand(todoAdd, "todo", "Add an item to the todo list (dev only)"))
	ch.cmds = append(ch.cmds, newCommand(todoDelete, "todo delete <i>", "Delete the item i from the list (dev only)"))
	ch.cmds = append(ch.cmds, newCommand(todoRead, "todo list", "List the current todo list (dev only)"))

	ch.generateHelpFunction()

	for _, c := range ch.cmds {
		ch.dg.AddHandler(c.function)
	}
}

//Dynamically generate the help function based on commands in the struct
func (ch commandHandler) generateHelpFunction() {
	//lambda help function added to the discordgo handler
	ch.dg.AddHandler(func(discord *discordgo.Session, message *discordgo.MessageCreate) {
		// Ignore all messages created by the bot itself
		if message.Author.ID == discord.State.User.ID {
			return
		}

		if message.Content == prefix + "help" {
			var embedFields []*discordgo.MessageEmbedField
			for _, c := range ch.cmds {
				embedFields = append(embedFields, &discordgo.MessageEmbedField{
					Name:   prefix + c.cmd,
					Value:  c.description,
					Inline: true,
				})
			}

			//Use embeds as the help
			embed := &discordgo.MessageEmbed{
				Author:      &discordgo.MessageEmbedAuthor{},
				Color:       0x696969, // Grey
				Description: "Contact thehowlinggreywolf#5036 if you're experiencing issues",
				Fields:      embedFields,
				//Timestamp: time.Now().Format(time.RFC3339), // Discord wants ISO8601; RFC3339 is an extension of ISO8601 and should be completely compatible.
				Title: "A List of Commands",
			}

			_, err := discord.ChannelMessageSendEmbed(message.ChannelID, embed)
			if err != nil {
				log.Fatal(err)
				return
			}
		}

	})
}
