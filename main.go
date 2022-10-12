package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const confrenceName = "GopherCon"
const maxAttendees = 50
var wg = sync.WaitGroup{}

type userData struct {
	name string
	email string
	age uint
	needSeats int
}
	

func main() {
	
	var remainingSeats uint = maxAttendees

	var bookings []userData = make([]userData, 0, maxAttendees);

	for {
		if !greetUsers(remainingSeats) {
			break
		}
		
		name, email, age, needSeats := getUsersInput()

		isValidName, isValidEmail, isValidSeats := validateUserInputs(name, email, needSeats, remainingSeats)

		if isValidName && isValidEmail && isValidSeats {
			attendent := bookTickits(name, email, age, needSeats, &remainingSeats, &bookings)
			wg.Add(1)
			go sendTickits(*attendent)
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
	wg.Wait()
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
	var scanner = bufio.NewScanner(os.Stdin)		
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

func bookTickits(name string, email string, age uint, needSeats int, remainingSeats *uint, bookings *[]userData) *userData {
	newUser := NewUser(name,email,age,needSeats)
	*remainingSeats -= uint(needSeats)
	*bookings = append(*bookings, *newUser)
	return newUser
}

func NewUser(name string, email string, age uint, needSeats int) *userData {
	user := new(userData)
	user.name = name
	user.age = age
	user.email = email
	user.needSeats = needSeats

	return user
}

func sendTickits(attendent userData){
	time.Sleep(20 * time.Second)
	var emailText = fmt.Sprintf("Hey %v,\n\tThanks for registration for %v\n Here are Your %v Tickits",attendent.name,confrenceName,attendent.needSeats)
	fmt.Println(emailText + "\n Send to " + attendent.email)
	wg.Done()
}