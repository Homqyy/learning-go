package main

import (
	"fmt"

	q "rsc.io/quote/v4"

	"example.com/greetings"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(q.Hello())

	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
