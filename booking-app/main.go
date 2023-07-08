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
	fmt.Println("Welcome to our", conferenceName, "application")
	fmt.Println("Get your tickets her to attend", "only ", remainingTickets, "tickets remaining")

	fmt.Println(conferenceName)
	for(remainingTickets > 0){
		remainingTickets--
		fmt.Println("Remaining tickets", remainingTickets)
	}
}
 