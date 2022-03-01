package time_mng

import (
	"sort"
	"time"
)

func CurrentDate() string {
	currentTime := time.Now()
	//return currentTime.Format("02-01-2006")
	return currentTime.Format("2006-01-02")
}

func DateToString(date time.Time) string {
	//return date.Format("02-01-2006")
	return date.Format("2006-01-02")
}

func SortStringDates(dates []string) ([]string, error) {

	var myTimeSlice timeSlice

	for _, date := range dates {
		newTime, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, err
		}
		myTimeSlice = append(myTimeSlice, newTime)
	}

	sort.Sort(timeSlice(myTimeSlice))

	var datesStringSorted []string

	for _, date := range myTimeSlice {
		datesStringSorted = append(datesStringSorted, DateToString(date))
	}

	return datesStringSorted, nil
}
