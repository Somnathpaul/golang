package main

import (
	"fmt"
	"reflect"
)

func main() {
	// logical operator: ! = not , && = and , || = or
	fmt.Println(5 > 10) // expected: false
	fmt.Println(5 > 4)  // expected true
	fmt.Println(9 / 2)
	fmt.Println("reminder: ", 9%2)
	fmt.Println(9.0 / 2)

	// find type of variable with reflect
	var n = true
	var s = "golang"
	var f = 2.999
	fmt.Println(reflect.TypeOf(n))
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(f))

	// complex numbers: i represent complex number in go
	fmt.Println("4n +2: ", 11+4i)
	var c = 11 + 3i
	fmt.Println(reflect.TypeOf(c)) // expected: complex128
}
