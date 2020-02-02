package main

import "fmt"

func main() {
	//for loop
	/*
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d \n", i)
		}
	*/

	// fizzbuzz chalange:
	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("bizz")
		} else {
			fmt.Println(i)
		}
	}
}
