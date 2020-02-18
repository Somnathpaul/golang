package main

import (
	"fmt"
)

func main() {
	// ASCII code: 'h', always in single code with single character.
	fmt.Println("h:", 'h')
	fmt.Println("H :", 'H')
	//fmt.Println('hq') // Error, always single character, no double character allowed
	fmt.Println('i')
	fmt.Println('j')

	fmt.Println(`The American Standard Code for Information Interchange, or ASCII code, was created in 1963 by the American Standards Association Committee or ASA, the agency changed its name in 1969 by American National Standards Institute or ANSI as it is known since.
		This code arises from reorder and expand the set of symbols and characters already used in telegraphy at that time by the Bell company.`)
}
