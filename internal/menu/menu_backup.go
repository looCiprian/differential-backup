package menu

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func menu_backup(currentText string) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--destination", Description: "Directory to store the backup"},
		{Text: "--source", Description: "Directory to backup"},
	}

	if strings.Contains(currentText,"--source"){
		s = []prompt.Suggest{
			{Text: "--destination", Description: "Directory to store the backup"},
		}
		return s
	}
	if strings.Contains(currentText,"--destination"){
		s = []prompt.Suggest{
			{Text: "--source", Description: "Directory to backup"},
		}
		return s
	}
	if strings.Contains(currentText,"--destination") && strings.Contains(currentText,"--source"){
		s = []prompt.Suggest{
		}
		return s
	}
	return s
}
