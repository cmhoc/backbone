package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf, err := loadConfig("config.toml")
	if err != nil {
		log.Fatal("Could not load config")
	}

	dg, err := discordgo.New("Bot " + conf.Bottoken)
	if err != nil {
		log.Fatal(err)
	}

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
	}

	comHandler := newCommandHandler(dg)
	comHandler.loadCommands()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	_ = dg.Close()

}
