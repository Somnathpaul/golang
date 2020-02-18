package main

import (
	"fmt"
	"reflect"
)

// Go support uint, int:
// uint contains only positive numbers from 0
// init contains positive and negative numbers
// uint8, uint16, uint32, uint64
// int8, int16, int32, int64
// uint8: from 0 to 155
// int16: -32768 to +32767
// int32: -2147483648 to +2147483647.
// int64:  -9223372036854775808 to +9223372036854775807

// global variable
const (
	i    uint8  = 255
	name        = "golang"
	x           = 100
	z    uint16 = 32767
)

func main() {

	// non global variable
	var j = "j"
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.TypeOf(j))

	fmt.Printf("Value: %#v %T\n", z, z)

	// call the custom function to work
	foo()
}

func foo() {
	// non global
	var b bool = true
	fmt.Println(x)
	fmt.Println(b)
}
