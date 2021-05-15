package time_mng

import (
	"sort"
	"time"
)

func CurrentDate() string {
	currentTime := time.Now()
	return currentTime.Format("02-01-2006")
}

func DateToString(date time.Time) string {
	return date.Format("02-01-2006")
}

func SortStringDates(dates []string) ([]string, error) {

	var myTimeSlice timeSlice

	for _, date := range dates{
		newTime, err := time.Parse("02-01-2006", date)
		if err != nil{
			return nil, err
		}
		myTimeSlice = append(myTimeSlice, newTime)
	}

	sort.Sort(timeSlice(myTimeSlice))

	var datesStringSorted []string

	for _, date := range myTimeSlice{
		datesStringSorted = append(datesStringSorted, DateToString(date))
	}

	return datesStringSorted,nil
}