package main

import (
	"encoding/gob"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strconv"
	"strings"
)

func todoRead(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, prefix+"todo list") {
		if message.Author.ID == "155084706868625408" {
			var todoList []string
			file, err := os.Open("todo.gob")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			stat, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			if stat.Size() != 0 {
				dec := gob.NewDecoder(file)
				err = dec.Decode(&todoList)
				if err != nil {
					log.Fatal(err)
				}
			}
			if len(todoList) != 0 {
				_, _ = discord.ChannelMessageSend(message.ChannelID, "Todo List:")
				for index, item := range todoList {
					_, _ = discord.ChannelMessageSend(message.ChannelID, fmt.Sprintf("%d: %s", index+1, item))
				}
			} else {
				_, _ = discord.ChannelMessageSend(message.ChannelID, "Todo List is Empty")
			}
		}
	}
}

func todoDelete(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, prefix+"todo delete") {
		if message.Author.ID == "155084706868625408" {
			var todoList []string
			file, err := os.Open("todo.gob")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			stat, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			if stat.Size() != 0 {
				dec := gob.NewDecoder(file)
				err = dec.Decode(&todoList)
				if err != nil {
					log.Fatal(err)
				}
			}
			index1 := strings.TrimPrefix(message.Content, prefix+"todo delete ")
			index, err := strconv.Atoi(index1)
			//dealing with errors
			if err != nil {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Index not given")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			if (index-1 > len(todoList)) || (index-1 < 0) {
				_, err := discord.ChannelMessageSend(message.ChannelID, "Index not in range")
				if err != nil {
					log.Fatal(err)
				}
				return
			}
			index--
			todoList = append(todoList[:index], todoList[index+1:]...)
			file, err = os.Create("todo.gob")
			if err != nil {
				log.Fatal(err)
			}
			enc := gob.NewEncoder(file)
			err = enc.Encode(todoList)
			if err != nil {
				log.Fatal(err)
			}

			_, _ = discord.ChannelMessageSend(message.ChannelID, "Content Removed")

		} else {
			return
		}
	}
}

func todoAdd(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if strings.Contains(message.Content, prefix+"todo") {
		if strings.Contains(message.Content, prefix+"todo list") {
			return
		} else if strings.Contains(message.Content, prefix+"todo delete") {
			return
		} else if strings.TrimPrefix(message.Content, prefix+"todo") == "" {
			return
		} else if message.Author.ID == "155084706868625408" { //Only recognize my input
			var todoList []string
			//trimming prefix to just get the message
			temp := strings.TrimPrefix(message.Content, prefix+"todo ")
			file, err := os.Open("todo.gob")
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			stat, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}
			if stat.Size() != 0 {
				dec := gob.NewDecoder(file)
				err = dec.Decode(&todoList)
				if err != nil {
					log.Fatal(err)
				}
			}
			todoList = append(todoList, temp)
			file, err = os.Create("todo.gob")
			if err != nil {
				log.Fatal(err)
			}
			enc := gob.NewEncoder(file)
			err = enc.Encode(todoList)
			if err != nil {
				log.Fatal(err)
			}

			_, _ = discord.ChannelMessageSend(message.ChannelID, "Content Added")

		} else {
			return
		}
	}
}
