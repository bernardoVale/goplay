package main

import "fmt"

type bytesize float64

const(
	_ = iota
	KB bytesize = 1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	// Lets say that this is 1byte
	var a int = 1

	fmt.Println("Byte:", a)
	fmt.Println("Kilobyte:", a<<10)
	fmt.Println("Megabyte:", a<<20)
	fmt.Println("Gigabyte:", a<<30)

	giga := a<<30

	fmt.Println("Number of Gigabytes:", giga>>30)
	fmt.Println("How much in Kylobytes:", KB*bytesize(a))
	fmt.Println("How much in Mebabytes:", MB*bytesize(a))
	fmt.Println("How much in Gigabytes:", GB*bytesize(a))
	fmt.Println("How much in Terabytes:", TB*bytesize(a))
}