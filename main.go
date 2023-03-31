package main

import (
	"fmt"

	"github.com/geronimosua/booking-app/helper"
)

func main() {
	helper.Greet_users()

	for helper.Remaining_tickets != 0 {
		user_name, user_last_name, user_email, user_tickets := helper.Get_user_inputs()

		if helper.Is_valid_input(user_tickets, user_name, user_last_name, user_email) {
			helper.Book_tickets( user_tickets, user_name, user_last_name, user_email)
			first_names := helper.Get_first_names()
			fmt.Printf("The first names of the bookings are %v\n", first_names)
		}
	}

	fmt.Printf("All tickets were sold, see you in the next year")
	helper.WG.Wait()
}

