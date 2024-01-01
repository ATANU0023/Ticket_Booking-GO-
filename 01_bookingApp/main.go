package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go conference" //u can explicitly declare the type
const conferenceTickets = 50

var remainingTickets uint = 50     //var remainingTickets uint = 50  for uint the value can never be negetive.
var bookings = make([]UserData, 0) //slice() //map type
//here we initializing a list of map

type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	// conferenceName := "Go conference" //u can declare like this also for only variable this and the next line is doing the same work.

	//fmt.Printf("the type of the conferenceName is %T , conferenceTicket is %T , AND remainingTickets is %T\n",conferenceName,conferenceTickets,remainingTickets);

	greetUsers()
	//for {
		//var bookings [50]string //in go array has fixed sizes(var variable_name [size]variable_type)
		//declering array string type to give explicit input (var variable-name []vartype{ //here give the input})

		//ask user for their name

		//fmt.Println(remainingTickets)  //this will show u the value that u have initialized but
		//fmt.Println(&remainingTickets) //this will show the memory location /this is pointer

		//userInput function
		firstName, lastName, email, userTickets := getUserInput()

		//here the function for validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		//isValidCity := city == "Singapore" || city == "London"

		if isValidName && isValidEmail && isValidTicketNumber {

			//booklogic function here

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) //run concurrent way on a different thread

			/*
				fmt.Printf("the whole slice : %v\n", bookings)
				fmt.Printf("firstvalue %v\n", bookings[0])
				fmt.Printf("slice types %T\n", bookings)
				fmt.Printf("slice length %v\n", len(bookings))
			*/

			//function that prints first name here !
			firstNames := getFirstNames()
			fmt.Printf("The first names of the bookings are %v\n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Printf("Our conference is booked out . Come back next year ")
				//break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short ")
			}
			if !isValidEmail {
				fmt.Println("email address you entered  not contain @ sign ")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid ")

			}
			fmt.Println("\t Try again")
			//fmt.Printf("We have only %v  Tickets remaining. so you cann't book %v tickets\n", remainingTickets, userTicket)

		}
	//}
	wg.Wait()

	//userName ="tom"
	//userTicket = 3
	//fmt.Printf("user %v booked %v tickets",firstName,userTicket)

	//go has only for loop support

}

func greetUsers() {
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("we have total of  %v  tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, bookings := range bookings { //Range iterates over elements for different data structure (so not only arrays and slice)
		//for arrays and slices , range provides the index and value for each element.
		//var names = strings.Fields(booking) //split the string with white spaces as separate and returns a slice with the split elements
		firstNames = append(firstNames, bookings.FirstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\nPlease enter your First Name: ")
	fmt.Scan(&firstName) //here the pointer will stop and will ask for the name/

	fmt.Println("\nPlease enter your Last Name: ")
	fmt.Scan(&lastName)

	fmt.Println("\nPlease enter your Email here(Must have '@' sign): ")
	fmt.Scan(&email)

	fmt.Println("\nPlease enter the NO. of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create map for user
	//var mySlice = []string;
	//var myMap = [string]string
	var userData = UserData{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		NumberOfTickets: userTickets,
	}

	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, userData) //more dynamic (slice)
	fmt.Printf("List of Bookings is: %v\n", bookings)
	fmt.Println(" ")

	fmt.Printf("Thank you! %v %v for booking %v Tickets, You will recive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v Tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) //to make delay for 10 second

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)

	fmt.Println(" _________________________")
	fmt.Printf("Sending  tickets: \n %v \n to email address: %v\n", ticket, email)
	fmt.Println(" _________________________")
	wg.Done()
}
