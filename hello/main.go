package main
import (
	"fmt"
	"greetings"
	"log"
  "rsc.io/quote"
)

func main() {
	log.SetPrefix("messages:")
	log.SetFlags(0)
	fmt.Println(quote.Go())


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
