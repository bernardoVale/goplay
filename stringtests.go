package main

import "fmt"

func main() {

	var s string = "this is a string"

	fmt.Println(s)

	s += " with addend"
	fmt.Println(s)

	s = "some string"+
	"\nsome other"
	fmt.Println(s)
}
