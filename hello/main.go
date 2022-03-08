package main
import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("messages:")
	log.SetFlags(0)


	names := []string{
		"Jack",
		"Tom",
	}
	message, err := greetings.Hellos(names)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(message)

}
