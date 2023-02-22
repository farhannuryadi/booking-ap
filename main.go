package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

// variable lavel package
var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = 50
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

// structs
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main(){
	// var conferenceName = "Go Conference" // membuat variable dan langsung diinisialisasikan
	// var conferencaName string = "Go Conference" // bisa cara ini
	// conferencaName := "Go Conference" // atau seperti ini
	// const conferenceTickets = 50 // membuat constanta dan langsung diinisialisasikan
	// const conferenceTickets string = 50 // bisa cara ini
	// conferenceTickets := 50 // untuk const tidak bisa seperti ini
	// var remainingTickets uint = 50

	// var bookings = [50]string //deklarasi array 
	// var bookings = []string // deklarasi slice
	// bookings := []string{} // deklarasi slice

	greetUsers()

	for {
		
		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValiduserTicket := helper.ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValiduserTicket {
			
			bookTicket(userTicket, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Printf("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValiduserTicket {
				fmt.Println("number of ticket you entered is invalid")
			}
		}
	}
	wg.Wait()
}

func greetUsers()  {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string{
	firstNames := []string{}
	for _, booking := range bookings{ // _ hanya menerima pengembalian data saja, namun variable nya tidak dibutuh kan
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string // deklarasi variable
	var lastName string
	var email string
	var userTicket uint 

	fmt.Println("Enter your first name : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address : ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets : ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string)  {
	remainingTickets -= userTicket

	// map 
	// var userData = make(map[string]string)
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTicket), 10)

	// structs
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list of bookings is %v\n",bookings)

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("Slice type: %T\n", bookings)
	// fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will a confirmation email at %v\n",firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceTickets)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string)  {
	time.Sleep(10 * time.Second)
	fmt.Println("##########################")
	var ticket = fmt.Sprintf("%v tickects fo %v %v", userTicket, firstName, lastName)
	fmt.Printf("Sending ticket:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("##########################")
	wg.Done()
}