package main

import (
	"fmt"
)

func add(x int) int{
	return x + 2

}

// When several arguments follow the same type you can omit and name it just once

func myfunc(x, y, z int) int{

	return x + y + z
}

func multi_result() (int, string) {

	return 2, "test"
}

// Named Function example it will return vars declared on function header
func named_func() (x , y string){
	x = "Hello"
	y = "World"
	return
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var x, y, z int = 1, 2, 3

	fmt.Println(add(2))

	fmt.Println("Result of 1+2+3 is ", myfunc(x, y, z))

	// Python style :)
	test, test2 := multi_result()

	fmt.Println(test, test2)

	fmt.Println(named_func())
}
