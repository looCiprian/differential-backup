package menu

import (
	"diff-backup/internal/db_mng"
	"diff-backup/internal/file_mng"
	"github.com/c-bata/go-prompt"
	"strings"
)

func menu_restore(currentText string, currentWord string) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "--source", Description: "<directory_containing_backups> Directory containing the backup"},
		{Text: "--destination", Description: "<directory_to_save_restored_files> Directory to store the restored backup"},
	}

	if strings.Contains(currentText, "--source") && !strings.Contains(currentText, "--destination"){
		s := []prompt.Suggest{
			{Text: "--destination", Description: "<directory_to_save_restored_files> Directory to store the restored backup"},
		}
		return s
	}

	if strings.Contains(currentText, "--source") && strings.Contains(currentText, "--destination"){

		if strings.Contains(currentWord, "--date") && strings.Contains(currentText, "--source") {

			command := strings.Fields(currentText)
			var dates []string

			for i, option := range command{
				if strings.Contains(option,"--source"){
					if len(command) > i{
						if command[i+1] != "--date" {
							err := db_mng.OpenDB(file_mng.AddSlashIfNotPresent(command[i+1]) + "index.db")
							if err == nil{
								dates,_ = db_mng.GetAvailableRestoreDates()
								db_mng.CloseDB()
								break
							}
						}
					}
				}
			}

			if len(dates) == 0{
				return []prompt.Suggest{
					{Text: "", Description: "No available backups!!!"},
				}
			}

			var s []prompt.Suggest

			for _, date := range dates{
				s = append(s, prompt.Suggest{Text: date, Description: "Restore: " + date})
			}

			return s
		}

		if strings.Contains(currentText, "--source") && strings.Contains(currentText, "--destination") && strings.Contains(currentText, "--date"){
			s := []prompt.Suggest{
			}
			return s
		}

		s := []prompt.Suggest{
			{Text: "--date", Description: "<start_date_to_restore> Backup date to restore"},
		}

		return s
	}

	if strings.Contains(currentText, "--destination"){
		s := []prompt.Suggest{
			{Text: "--source", Description: "<directory_containing_backups> Directory containing the backup"},
		}
		return s
	}


	if strings.Contains(currentText, "--date"){
		s := []prompt.Suggest{
			{Text: "--source", Description: "<directory_containing_backups> Directory containing the backup"},
			{Text: "--destination", Description: "<directory_to_save_restored_files> Directory to store the restored backup"},
		}
		return s
	}

	return s
}
