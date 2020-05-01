package main

import (
	"fmt"
	"log"
	"os"
)

const englishHelloPrefix = "Hello, "

// Hello : An abstracted function, whichgo  is easier to test
func Hello(name string) string {
	return englishHelloPrefix + name + "!"
}

func main() {

	if len(os.Args) > 1 {
		name := os.Args[1]
		fmt.Println(Hello(name))
	} else {
		log.Print("Name argument not provided, using default.")
		fmt.Println(Hello("Apurv"))
	}

}
