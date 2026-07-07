package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "aremyparentsproudofme",
		Description: "Ask and find out!",
	},
	{
		Name:        "info",
		Description: "Information about the bot!",
	},
	{
		Name:        "ben",
		Description: "Ben.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "question",
				Description: "Question you want to ask Ben",
				Required:    true,
			},
		},
	},
	{
		Name:        "kashira",
		Description: "I suppose bro",
	},
}

func main() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent

	discord.AddHandler(messageCreate)
	discord.AddHandler(interactionCreate)

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

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, "", v)
		if err != nil {
			fmt.Println(err)
		}
		registeredCommands[i] = cmd
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "Skibidi Rizz")
		if err != nil {
			return
		}
	}

	if m.Content == "!isthisablunder" {
		message := &discordgo.MessageSend{
			Content:   chessRatings(),
			Reference: m.MessageReference,
		}
		_, err := s.ChannelMessageSendComplex(m.ChannelID, message)
		if err != nil {
			return
		}
	}
}

func interactionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionApplicationCommand {
		commandData := i.ApplicationCommandData()

		if commandData.Name == "info" {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "(this description will change later) \n Hello my name is Nieśmieszny your favorite slop YouTuber, I used to make shitty daily animations about stuff with no research, also I once accidently promoted drugs to children, oh and im doing Minecraft now",
				},
			})
			if err != nil {
				fmt.Println(err)
			}
		}

		if commandData.Name == "aremyparentsproudofme" {
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "No.",
				},
			})
			if err != nil {
				fmt.Println(err)
			}
		}

		if commandData.Name == "ben" {
			question := commandData.Options[0].StringValue()
			reply := benQuestion()
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Your question was: " + question + "\n\n" + reply,
				},
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(question)
		}

		if commandData.Name == "kashira" {
			filepaths := []string{
				"bettyrezero/angry.jpg",
				"bettyrezero/aura.jpg",
				"bettyrezero/goober.jpg",
				"bettyrezero/huh.jpg",
				"bettyrezero/kashirabro.jpg",
				"bettyrezero/pout.jpg",
				"bettyrezero/what.jpg",
			}
			filenames := []string{
				"angry.jpg",
				"aura.jpg",
				"goober.jpg",
				"huh.jpg",
				"kashirabro.jpg",
				"pout.jpg",
				"what.jpg",
			}
			imageID := rand.Intn(len(filenames))

			bettyImage, cleanup, err := readImage(filepaths[imageID], filenames[imageID])
			if err != nil {
				fmt.Println(err)
			}
			defer cleanup()

			err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Files: []*discordgo.File{bettyImage},
				},
			})
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// Functions outside discordgo but used for the commands

func benQuestion() string {
	responses := []string{
		"Yes",
		"No",
		"Hohoho",
		"Ughhhh",
	}
	return responses[rand.Intn(len(responses))]
}

func chessRatings() string {
	ratings := []string{
		"<:Blunder:1517582398374674582>",
		"<:Blunder:1517582398374674582>",
		"<:Best:1520800747246063746>",
		"<:Best:1520800747246063746>",
		"<:Mistake:1520800166204936243>",
		"<:Mistake:1520800166204936243>",
		"<:Excellent:1521240041823666306>",
		"<:Excellent:1521240041823666306>",
		"<:Bookmove:1520799464057733373>",
		"<:Bookmove:1520799464057733373>",
		"<:Great:1520800903504990339>",
		"<:Great:1520800903504990339>",
		"<:Brilliant:1517582607448281198>",
	}
	return ratings[rand.Intn(len(ratings))]
}

func readImage(path string, fileName string) (*discordgo.File, func(), error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	bettyImage := &discordgo.File{
		Name:   fileName,
		Reader: file,
	}
	cleanup := func() {
		_ = file.Close()
	}
	return bettyImage, cleanup, nil
}
