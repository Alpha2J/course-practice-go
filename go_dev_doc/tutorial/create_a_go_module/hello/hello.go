package main

import (
	"fmt"
	"hellooo.site/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)
	//message, err := greetings.Hello("Gladys")
	// uncomment me to see errors
	//message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(message)
	fmt.Println(messages)
}
