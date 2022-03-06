package main
import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("messages:")
	log.SetFlags(0)


	message, err := greetings.Hello("TOm")
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)

}