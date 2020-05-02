package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/apurvmishra99/integers"
)

var calcOperations = map[string]func(int, int) int{
	"+": integers.Add,
	"-": integers.Subtract,
	"*": integers.Multiply,
	"/": integers.Division}

func calculator(operation string, x, y int) int {
	return calcOperations[operation](x, y)
}

func main() {
	if len(os.Args) > 3 {
		op := os.Args[1]
		x, err1 := strconv.Atoi(os.Args[2])
		y, err2 := strconv.Atoi(os.Args[3])

		if err1 == nil && err2 == nil {
			if _, ok := calcOperations[op]; ok {
				fmt.Println(calculator(op, x, y))
			} else {
				fmt.Println("Only '+', '-', '*' and '/' are permitted.")
			}
		} else {
			fmt.Println("Provide integral arguments.")
		}

	} else {
		fmt.Println("Usage : op : [+,-,*,/] x : int y : int")
	}

}
