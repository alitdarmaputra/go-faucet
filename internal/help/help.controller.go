package help

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type HelpHandler struct {
	helpService *HelpService
}

func NewHelpHandler(helpService *HelpService) *HelpHandler {
	return &HelpHandler{
		helpService: helpService,
	}
}

func (h *HelpHandler) Handler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content == "!h" {
		log.Printf("info: Received help operation from %s - (Channel: %s)\n", m.Author.Username, m.ChannelID)
		var content string = h.helpService.HelpService(m)
		_, err := s.ChannelMessageSend(m.ChannelID, content)
		if err != nil {
			fmt.Println("Error:", err)
		}
		return
	}

	if strings.HasPrefix(m.Content, "!usage") {
		log.Printf("info: Received help usage operation from %s - (Channel: %s)\n", m.Author.Username, m.ChannelID)
		// Split argument
		args := strings.Split(m.Content, " ")

		if len(args) == 2 {
			content := h.helpService.GetHelp(args[1])
			_, err := s.ChannelMessageSend(m.ChannelID, content)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		_, err := s.ChannelMessageSend(
			m.ChannelID,
			"Not enough argument. Please provide command name",
		)
		if err != nil {
			log.Println("error: sending discord message\n", err)
		}
	}
}
