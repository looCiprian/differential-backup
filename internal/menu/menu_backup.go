package menu

import (
	"github.com/c-bata/go-prompt"
	"strings"
)

func menuBackup(currentText string) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "--destination", Description: "<initialized_backup_dir> Directory to store the backup"},
		{Text: "--source", Description: "<directory_to_backup> Directory to backup"},
	}

	if strings.Contains(currentText, "--source") {
		s = []prompt.Suggest{
			{Text: "--destination", Description: "<initialized_backup_dir> Directory to store the backup"},
		}
		return s
	}
	if strings.Contains(currentText, "--destination") {
		s = []prompt.Suggest{
			{Text: "--source", Description: "<directory_to_backup> Directory to backup"},
		}
		return s
	}
	if strings.Contains(currentText, "--destination") && strings.Contains(currentText, "--source") {
		s = []prompt.Suggest{
		}
		return s
	}
	return s
}
