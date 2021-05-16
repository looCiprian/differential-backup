package menu

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func menuInit(currentText string) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--destination", Description: "<init_backup_dir> Set the repository"},
	}

	if strings.Contains(currentText, "--destination") {
		s = []prompt.Suggest{
		}
		return s
	}
	return s
}
