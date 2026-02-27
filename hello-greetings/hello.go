package main

import (
	"fmt"
	"log"
	"coelho.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string {
		"Aron",
		"Mathew",
		"Tom",
	}
	// message, err := greetings.Hello("Aron")
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}