package main

import (
	"fmt"

	"github.com/FeawBook/cinema/movie"
	"github.com/FeawBook/cinema/ticket"
)

func main() {
	movieName := movie.FindName("1234")
	fmt.Println(movieName)
	movie.Review(movieName, 10.0)
	ticket.BuyTicket(movieName)
}
