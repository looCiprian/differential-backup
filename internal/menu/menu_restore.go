package menu

import (
	"github.com/c-bata/go-prompt"
)

func menu_restore() []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--source", Description: "Directory containing the backup"},
		{Text: "--destination", Description: "Directory to store the retored backup"},
	}
	return s
}
