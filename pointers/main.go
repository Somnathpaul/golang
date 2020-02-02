package main

import "fmt"

func main() {
	i, j := 2, 80
	// initialise pointer p to i
	p := &i
	// print the value of i with p
	fmt.Println(*p)
	// set new value of i with p
	*p = 21
	fmt.Printf("the new value is : %d \n", *p)

	fmt.Println(j)
	// initilise pointer p to j
	p = &j // already assigned above, so : not required
	// divide the value of j by 8 with pointer p
	*p = *p / 80
	// print the new value of j
	fmt.Printf("the new value of j %d \n", *p)
}
