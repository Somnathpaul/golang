package main

import "fmt"

func main() {
	// arrays : will have fixed length

	// delare and assign value:
	arra1 := [2]string{"orange", "apple"}

	// slices: fixed length not required

	// declare and assign:
	slice1 := []string{"orange", "apple", "banana", "watermellon"}

	//print
	fmt.Println(arra1)
	fmt.Println(slice1)

	//length of an slice
	fmt.Println(len(slice1))

	s := []int{2, 3, 4, 5, 6, 7}
	fmt.Println(s)

	// slice an array or slice
	fmt.Println(slice1[1:3])

	// define a slice with make
	a := make([]int, 2, 5) // length and capacity
	a[0] = 10
	a[1] = 20
	a[3] = 30
	fmt.Println("Slice A:", a)

	// define a slice with new
}
