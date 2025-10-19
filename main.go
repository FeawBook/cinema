package main

import (
	"fmt"

	"github.com/BookFeaw/cinema/movie"
	"github.com/BookFeaw/cinema/ticket"
)

func main() {
	movieName := movie.FindName("1234")
	fmt.Println(movieName)
	movie.Review(movieName, 10.0)
	ticket.BuyTicket(movieName)
}
