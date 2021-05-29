package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pagenumber uint
	var grade float32
	title = "Mr. GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "DizzyBooks Publishing Inc."
	pagenumber = 14
	year = 1997
	grade = 6.5
	fmt.Println(title, " written by ", writer, " drawn by ", artist, " Published by ", publisher)
	fmt.Println("The Comic has ", pagenumber, " pages and came out in ", year, ". It is currently valued at ", grade)

	title = "Epic Vol. 1"
	writer = "Ryan N. Shawn"
	artist = "Phoebe Paperclips"
	publisher = "DizzyBooks Publishing Inc."
	pagenumber = 160
	year = 2013
	grade = 9.0
	fmt.Println(title, " written by ", writer, " drawn by ", artist, " Published by ", publisher)
	fmt.Println("The Comic has ", pagenumber, " pages and came out in ", year, ". It is currently valued at ", grade)
}
