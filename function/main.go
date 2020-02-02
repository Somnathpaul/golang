package main

import (
	"fmt"

	"github.com/somnath/go_crashCourse/packages/divide"
	"github.com/somnath/go_crashCourse/packages/multiply"
)

// custom method
func greetings(name string) string {
	return "Hello " + name
}

// custom method
func add(num1 int, num2 int) int {
	return num1 + num2
}

// main method
func main() {
	fmt.Println(greetings("somnath"))
	fmt.Println(add(2, 3))
	fmt.Println(multiply.Multiply(5, 5))
	fmt.Println(divide.Divide(90, 9))
}
