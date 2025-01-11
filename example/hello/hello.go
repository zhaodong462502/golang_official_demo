package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// fmt.Println("Hello World!")
	// fmt.Println(quote.Go())

	//message, err := greetings.Hello("Gladys")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(message)
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
