package main

import "fmt"

func main() {
	// conferenceName := "GoConf"
	// const numberOfTickets = 100
	// remainingTickets := 100

	// fmt.Printf("Welcome to %v event!\n", conferenceName)
	// fmt.Printf("There is %v avaliable ticket(s) from %v tickets.\n", remainingTickets, numberOfTickets)

	// var name string
	// var ticketCount int

	// fmt.Println("Enter your name: ")x``
	// fmt.Scan(&name)


	// fmt.Println("Enter ticket count: ")
	// fmt.Scan(&ticketCount)

	// fmt.Printf("> %v bought %v tickets.\n", name, ticketCount)

	// remainingTickets = remainingTickets - ticketCount

	// fmt.Printf("%v tickets remaining in conference!\n", remainingTickets)

	// array size is static
	// var users = [3] string {"Hamid", "Amir", "Majid"}
	// var users [3] string

	// slice size is dynamic
	// var users[]string
	// users = append(users, "Hamid")
	// users = append(users, "Amir")

	// fmt.Printf("Users are here: %v\n", users)

	var users []string

	for {
		var name string
		fmt.Println("Enter user or `exit`: ")
		fmt.Scan(&name)
		if (name == "exit") {
			break
		}
		users = append(users, name)
	}

	for _, name := range users {
		fmt.Println(name)
	}

	bye("Hamid")
}


func bye(name string) {
	fmt.Printf("Bye %v\n!", name)
}
