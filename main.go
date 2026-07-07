package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(discord)
	fmt.Println("Nieśmieszny jest online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Skibidi rizz")
		if err != nil {
			return
		}
	}
}
