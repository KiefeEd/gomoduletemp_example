package main

import (
	gotime "github.com/KiefeEd/gomoduletemp"
)

func main() {
	go gotime.DisplayTimeLocal() // displays local time based on where your server is hosted.
	go gotime.DisplayTime(7) //shows GMT+7 time now.
	go gotime.Greet("Tom", 4) // Greets user Tom who has a timezone of GMT+4.
	go gotime.GreetLocal("Tim") //Greets Tim based on local time.
}