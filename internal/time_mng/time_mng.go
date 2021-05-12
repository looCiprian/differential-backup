package time_mng

import "time"

func CurrentDate() string{
	currentTime := time.Now()
	return currentTime.Format("02-01-2006")
}
