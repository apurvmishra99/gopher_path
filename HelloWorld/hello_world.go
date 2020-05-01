package main

import "fmt"

// Hello : An abstracted function, whichgo  is easier to test
func Hello(name string) string {
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("Apurv"))
}
