package models


// Passenger schema of the passenger table
type Passenger struct {
	Passenger_id int64  `json:"id"`
	First_name   string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Phone int64  `json:"phone"`
	Email string `json:"email"`
	City  string `json:"city"`
}

type Ticket struct {
	Ticket_id int64 `json:"id"`
	Date      int64 `json:"date"`
}