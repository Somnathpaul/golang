package main

import (
	"fmt"
)

var val = 0

func main() {
	k := 1
	for ; k <= 15; k++ {
		fmt.Println("golang", val+1)
		val++
	}
}
