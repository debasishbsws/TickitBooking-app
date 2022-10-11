package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

	const confrenceName = "GopherCon"
	const maxAttendees = 50
	var scanner = bufio.NewScanner(os.Stdin)

func main() {
	
	var remainingSeats uint = maxAttendees

	var bookings []string = make([]string, 0, maxAttendees);

	for {
		if !greetUsers(remainingSeats) {
			break
		}
		
		name, email, age, needSeats := getUsersInput()

		isValidName, isValidEmail, isValidSeats := validateUserInputs(name, email, needSeats, remainingSeats)

		if isValidName && isValidEmail && isValidSeats {
			bookingMessage := bookTickits(name, email, age, needSeats, &remainingSeats, &bookings)
			fmt.Printf("%v Thanks for registration\n\n",bookingMessage)
		}else {
			if !isValidSeats {
				fmt.Printf("Invalid Tickits request. %d seats remaining.., please enter the appropreate number.\n\n", remainingSeats)
			}
			if !isValidName {
				fmt.Println("Invalid Name!!!")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email!!!")
			}
			fmt.Print("Please try again...\n\n")
		}
		
	}

	fmt.Println(bookings)


}

func greetUsers(remainingSeats uint) bool {
	fmt.Printf("Wellcome to %s booking application.\n\n", confrenceName)
		if remainingSeats == 0 {
			fmt.Printf("We are full!!, No seats available. Please come back next Year.\n")
			return false
		} else if remainingSeats < 5 {
			fmt.Printf("Only %d seats left! Book Your tickits now.\n", remainingSeats)
		} else {
			fmt.Printf("We have %d seates remaine.\n", remainingSeats)
		}
		fmt.Println("To Book Your Seats..")

		return true
}

func validateUserInputs(name, email string, needSeats int, remainingSeats uint) (bool, bool, bool) {
	isValidName := strings.Contains(name, " ")
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidSeats := needSeats >= 0 && (uint(needSeats) <= remainingSeats)
	return isValidName , isValidEmail, isValidSeats
}

func getUsersInput()(string, string, uint,int) {
	var (
			name 	string
			email 	string
			age uint
			needSeats int
		)
		fmt.Print("\nEnter Your Full Name : ")
		scanner.Scan()
		name = scanner.Text()

		fmt.Print("\nEnter your Age : ")
		fmt.Scan(&age)

		fmt.Print("\nEnter your Email : ")
		scanner.Scan()
		email = scanner.Text()

		fmt.Print("\nEnter How many seats you need : ")
		fmt.Scan(&needSeats)
	return name, email, age, needSeats
}

func bookTickits(name string, email string, age uint, needSeats int, remainingSeats *uint, bookings *[]string) string {
	bookingmessage := fmt.Sprintf("Your booking info\n\tConference Name: %s,\n\tTickit holder Name: %s,Age: %d,\n\tEmail: %s,\n\tTicket: %d\n", confrenceName, name, age, email, needSeats)
	*remainingSeats -= uint(needSeats)
	*bookings = append(*bookings, bookingmessage)
	return bookingmessage
}