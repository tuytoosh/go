package main

import "fmt"

func main() {
	conferenceName := "GoConf"
	const numberOfTickets = 100
	remainingTickets := 100
    fmt.Printf("Welcome to %v event!\n", conferenceName)
	fmt.Printf("There is %v avaliable ticket(s) from %v tickets.\n", remainingTickets, numberOfTickets)

	var name string
	var ticketCount int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)


	fmt.Println("Enter ticket count: ")
	fmt.Scan(&ticketCount)

	fmt.Printf("> %v bought %v tickets.\n", name, ticketCount)

	remainingTickets = remainingTickets - ticketCount

	fmt.Printf("%v tickets remaining in conference!\n", remainingTickets);
}
