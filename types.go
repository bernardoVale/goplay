package main

import "fmt"

type capa int

func main() {
	s := 2
	if capa(s) == 2{
		fmt.Printf("Cast funcionou")
	}
}
