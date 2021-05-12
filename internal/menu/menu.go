package menu

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "init", Description: "Init new backup repo"},
		{Text: "backup", Description: "Backup repo"},
		{Text: "restore", Description: "Restore a complete backup"},
		{Text: "exit", Description: "Exit"},
	}

	currentText := d.TextBeforeCursor()

	if strings.Contains(currentText, "init") ||  strings.Contains(currentText, "backup") ||  strings.Contains(currentText, "restore") || strings.Contains(currentText, "exit"){
		switch  strings.Fields(currentText)[0]{
		case "init":
			s=menu_init(currentText)
		case "backup":
			s=menu_backup(currentText)
		case "restore":
			s=menu_restore()
		case "exit":
			s=menu_exit()
		}
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func Base(message string) string {
	fmt.Println(message)
	t := prompt.Input("> ", completer)
	return t
}
