package main

import (
	"fmt"
	"time"
)

func task2(dobStr string) {
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %v\n", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday time.Weekday) {
	weekday = date.Weekday()
	return
}
