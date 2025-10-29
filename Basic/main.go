package main

import (
	"fmt"
)

func main() {

	fmt.Println("This is  GO Programming")

	// Variable in Go Lang

	// Method 1 (Can Be Used Inside Function Only)
	name := "Kunal Kumar"

	// Method 2 (Declared then Initialized)
	var regno string
	regno = "22MEI100399"

	// Method 3 (Decalare and Initialize)
	var school string = "SCOPE"

	var a, b, c int = 1, 2, 3

	// Default Value of Int is "0"
	var zero int

	fmt.Println(name, " ", regno, " ", school)
	fmt.Println(a, b, c, zero)

	// Print Type of Variable
	fmt.Printf("%T\n", name)

	const counter = 999999
	fmt.Println(counter)

	// Loops in Go Lang (Onlhy For Loop)
	i := 10
	for i != 0 {
		fmt.Println(i)
		i--
	}

	for j := 1; j <= 10; j++ {
		fmt.Println("Loop Count", j)
	}

}
