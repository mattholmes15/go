package main

import (
	"example/greetings"
	"fmt"
	"log"
	// "rsc.io/quote"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(3)

	message, err := greetings.Hello("Holmes")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Hello World")
	// fmt.Println(quote.Go())
	// message := greetings.Hello("Holmes")

	fmt.Println(message)
}
