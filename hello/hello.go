package main

import (
	"fmt"
	"kj-tech.net/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"KJ", "Jhon", "George"}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
