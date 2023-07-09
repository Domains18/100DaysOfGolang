package main

import "fmt"

func main() {
	//declaring variables
	/**
	@var - mutable
	@const immutable
	*/
	var conferenceName = " Go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	//println ensures next code printed out is on the next line
	//ptintf allows us to format the string
	fmt.Printf("Welcome to our %v application", conferenceName + "\n")
	fmt.Println("Get your tickets her to attend", "only ", remainingTickets, "tickets remaining")
	fmt.Printf("conference tickets are %v and remaining tickets are %v", conferenceTickets, remainingTickets)

}
 