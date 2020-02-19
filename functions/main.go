package main

import (
	"fmt"
)

func main() {

	// output store the return value and call the function
	var output = printLine(11)
	// print the return value
	fmt.Println(output)

	//store two return type in two different variable
	var swap1, swap2 = swap(10, 20)
	fmt.Println("After swap", swap1, swap2)
}

func printMessage() {
	for i := 0; i <= 10; i++ {
		fmt.Println("i")
	}
}

// function accept parameter along with the type
func printLine(val int) string {
	for i := 0; i <= val; i++ {
		fmt.Println(i)
	}
	return "done"
}

// swap function
func swap(a, b int) (int, int) {
	fmt.Println("Before swap", a, b)
	var c = a
	a = b
	b = c
	return a, b
}
