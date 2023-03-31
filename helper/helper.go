package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const Conference_name string = "Go Con";
const Conference_tickets uint8 = 50
var Remaining_tickets uint8 = 50
var Bookings = make([]UserData, 0)

var WG = sync.WaitGroup{}

type UserData struct {
	first_name string
	last_name string
	email string
	ticketsNumber uint8
}

func Greet_users() {
	fmt.Printf("Welcome to %v booking application\n", Conference_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", Conference_tickets, Remaining_tickets)
	fmt.Println("Get your tickets here to attend")
}

func Book_tickets(user_tickets uint8, user_name, user_last_name, user_email string) {
	Remaining_tickets -= user_tickets
	
	var userData = UserData {
		first_name: user_name,
		last_name: user_last_name,
		email: user_email,
		ticketsNumber: user_tickets,
	}
	
	Bookings = append(Bookings, userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", user_name, user_last_name, user_tickets, user_email)
	fmt.Printf("%v tickets remaining for %v\n", Remaining_tickets, Conference_name)
	WG.Add(1)
	go send_ticket_email(user_tickets, user_name, user_last_name, user_email)
	fmt.Printf("The booking list is now %v\n", Bookings)
}

func Get_user_inputs() (string, string, string, uint8) {
	var user_name string
	var user_last_name string
	var user_email string
	var user_tickets uint8

	fmt.Print("Enter your name: ")
	fmt.Scan(&user_name)
	
	fmt.Print("Enter your last name: ")
	fmt.Scan(&user_last_name)

	fmt.Print("Enter your email address: ")
	fmt.Scan(&user_email)
	
	fmt.Print("Enter your tickets: ")
	fmt.Scan(&user_tickets)
	return user_name, user_last_name, user_email, user_tickets
}

func Is_valid_input(user_tickets uint8, user_name, user_last_name, user_email string) bool {
	is_valid_name := len(user_name) >= 2 && len(user_last_name) >= 2
	is_valid_email := strings.Contains(user_email, "@")
	is_valid_ticket_number := user_tickets > 0 && user_tickets <= Remaining_tickets

	if !is_valid_email || !is_valid_name || !is_valid_ticket_number {
		if !is_valid_name {
			fmt.Printf("The name and the last name has to be at least 2 characters long\n")
		}
		if !is_valid_email {
			fmt.Printf("The email address doesn't contain a @ sign\n")
		}
		if !is_valid_ticket_number {
			fmt.Printf("You have to put a valid ticket number, remember it has to be at least more than 0 and that there are only %v tickets available\n", Remaining_tickets)
		}
		fmt.Printf("Pleasy try again\n")

		return false
	}
	
	return true
}

func Get_first_names() []string {
	first_names := []string{}

	for _, booking := range Bookings {
		first_names = append(first_names, booking.first_name)
	}
	return first_names
}

func send_ticket_email(user_tickets uint8, user_name, user_last_name, user_email string) {
	time.Sleep(10 * time.Second)
	message := fmt.Sprintf("%v tickets for %v %v", user_tickets, user_name, user_last_name)
	fmt.Println("#########")
	fmt.Printf("Sending email:\n %v to email address: %v\n", message, user_email)
	fmt.Println("#########")
	WG.Done()
}