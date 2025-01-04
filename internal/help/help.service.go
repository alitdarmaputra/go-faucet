package help

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type HelpService struct{}

func NewHelpService() *HelpService {
	return &HelpService{}
}

func (h *HelpService) HelpService(m *discordgo.MessageCreate) string {
	content := fmt.Sprintf("Hello %s, \n**Keywords: transfer**\n\nFor help type:\n!usage <available keyword>", m.Author.Username)
	return content
}

func (h *HelpService) GetHelp(command string) string {
	if command == "transfer" {
		return "Type:\n!transfer <wallet public key>\n**Ex: !transfer 0x... **"
	} else {
		return "Sorry, keyword not found"
	}
}
