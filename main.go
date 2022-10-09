package main

import (
	"bufio"
	"fmt"
	"os"
)


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	confrenceName := "GopherCon"
	const maxAttendees = 50
	var remainingSeats uint = maxAttendees

	var bookings []string = make([]string, 0, maxAttendees);

	for {
		fmt.Printf("Wellcome to %s booking application.\n\n", confrenceName)
		if remainingSeats == 0 {
			fmt.Printf("We are full!!, No seats available. Please come back next Year.\n")
			break
		} else if remainingSeats < 5 {
			fmt.Printf("Only %d seats left! Book Your tickits now.\n", remainingSeats)
		} else {
			fmt.Printf("We have %d seates remaine.\n", remainingSeats)
		}
		fmt.Println("To Book Your Seats..")

		var (
			name 	string
			email 	string
			age uint
			needSeats uint
		)
		fmt.Print("\nEnter Your Name : ")
		scanner.Scan()
		name = scanner.Text()

		fmt.Print("\nEnter your Age : ")
		fmt.Scan(&age)

		fmt.Print("\nEnter your Email : ")
		scanner.Scan()
		email = scanner.Text()

		fmt.Print("\nEnter How many seats you need : ")
		fmt.Scan(&needSeats)

		if needSeats > remainingSeats {
			fmt.Println("Sorry!! We can't book your tickits, Only", remainingSeats, "seats remain")
			fmt.Print("Please try again...\n\n")
			continue
		} else {
			bookingmessage := fmt.Sprintf("Your booking info\n\tConference Name: %s,\n\tTickit holder Name: %s,Age: %d,\n\tEmail: %s,\n\tTicket: %d\n", confrenceName, name, age, email, needSeats)
			remainingSeats -= needSeats
			bookings = append(bookings, bookingmessage)
			fmt.Printf("%v Thanks for registration\n\n",bookingmessage)
		}
	}

	fmt.Println(bookings)


}