package main

import (
	"fmt"
	"log"
	"os"
)

const englishHelloPrefix = "Hello, "

var greetings = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"Hindi":   "Namaste, "}

// Hello : An abstracted function, which is easier to test
func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetings[lang] + name + "!"
}

func main() {

	if len(os.Args) > 2 {
		name := os.Args[1]
		lang := os.Args[2]
		if _, ok := greetings[lang]; !ok {
			lang = "English"
		}
		fmt.Println(Hello(name, lang))
	} else {
		log.Print("Some arguments not provided, using default.")
		fmt.Println(Hello("Apurv", "English"))
	}

}
