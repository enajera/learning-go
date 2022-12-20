package main

import (
	"fmt"
	"github.com/enajera/learning-go/greet"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.Italian())
	fmt.Println(quote.Hello())
	fmt.Println(quoteV3.Concurrency())
	
}