package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/alitdarmaputra/go-faucet/config"
	"github.com/alitdarmaputra/go-faucet/internal/help"
	"github.com/alitdarmaputra/go-faucet/internal/transfer"
	"github.com/alitdarmaputra/go-faucet/pkg/ether"
	"github.com/bwmarrin/discordgo"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error reading .env file")
	}
}

func main() {
	config := config.GetConfig()

	// Initialize eth client
	client, err := ethclient.Dial(config.RpcUrl)
	if err != nil {
		log.Fatalln("error: Error connecting to eth client", err)
		return
	}
	defer client.Close()
	log.Println("info: Success connecting to eth client")

	// Opening discord session
	dg, err := discordgo.New("Bot " + config.BotToken)
	if err != nil {
		log.Fatalln("error: Error creating discord session:\n", err)
		return
	}

	// Registering handler
	ether := ether.NewEther(client)
	transferService := transfer.NewTransferService(config, ether)
	helperService := help.NewHelpService()

	dg.AddHandler(transfer.NewTransferHandler(transferService).Handler)
	dg.AddHandler(help.NewHelpHandler(helperService).Handler)
	log.Println("info: Success registering bot handler")

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatalln("error: Error opening discord session:\n", err)
		return
	}

	defer dg.Close()
	log.Println("info: Success opening discord session")

	// Wait until CTRL-C or other term signal is received
	log.Println("info: Go Faucet BOT is running...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}
