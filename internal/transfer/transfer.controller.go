package transfer

import (
	"context"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type TransferHandler struct {
	transferService *TransferService
}

func NewTransferHandler(transferService *TransferService) *TransferHandler {
	return &TransferHandler{
		transferService: transferService,
	}
}

func (t *TransferHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if strings.HasPrefix(m.Content, "!transfer") {
		log.Printf("info: Received transfer operation from %s - (Channel: %s)\n", m.Author.Username, m.ChannelID)
		var content string = ""

		args := strings.Split(m.Content, " ")

		if len(args) < 2 {
			content = "Please provide your wallet public key"
		} else {
			content = t.transferService.TransferService(context.Background(), args[1])
		}

		_, err := s.ChannelMessageSend(m.ChannelID, content)
		if err != nil {
			log.Println("error: sending discord message\n", err)
		}
		return
	}
}
