package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("The time is", currentTime) // 2024-04-20 17:34:35.5557453 +0100 BST m=+0.003327401
	fmt.Println(currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Hour(), currentTime.Minute(), currentTime.Second()) //2024 April 20 17 38 2
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime) //The time is 2021-08-15 14:30:45.0000001 +0100 BST
	fmt.Println(theTime.Format("2006-1-2 15:4:5")) //2021-8-15 14:30:45
	timeString := "2021-08-15 02:30:45"
	theTime, err := time.Parse("2006-01-02 03:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
	}
	fmt.Println("The time is", theTime) //The time is 2021-08-15 02:30:45 +0000 UTC

	toAdd := -24 * time.Hour
	fmt.Println("Adding", toAdd) //Adding -24h0m0s

	newTime := theTime.Add(toAdd)
	fmt.Println("The new time is", newTime) //The new time is 2021-08-14 02:30:45 +0000 UTC

	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime) // The first time is 2021-08-15 14:30:45.0000001 +0000 UTC

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime) // The second time is 2021-12-25 16:40:55.0000002 +0000 UTC

	fmt.Println("First time before second?", firstTime.Before(secondTime)) // First time before second? true
	fmt.Println("First time after second?", firstTime.After(secondTime)) // First time after second? false

	fmt.Println("Second time before first?", secondTime.Before(firstTime)) // Second time before first? false
	fmt.Println("Second time after first?", secondTime.After(firstTime)) // Second time after first? true

	fmt.Println("Duration between first and second time is", firstTime.Sub(secondTime)) //Duration between first and second time is -3170h10m10.0000001s
	fmt.Println("Duration between second and first time is", secondTime.Sub(firstTime)) //Duration between second and first time is 3170h10m10.0000001s

}
