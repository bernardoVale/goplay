package main

import "fmt"

var (
	be bool = true
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, be, be)

}