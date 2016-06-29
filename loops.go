package main

import "fmt"

func c_like() {
	sum := 0
	for i:=0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func while_like(){
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func main() {
	sum := 1
	for ; sum < 100; {
		sum += sum
	}
	fmt.Println(sum)
	while_like()
	c_like()
}
