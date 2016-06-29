package main

import "fmt"

const (
	Janeiro = iota // iota retorna 0 e incrementa
	Fevereiro
	Março
	Abril
	Maio
)
func main() {
	fmt.Println(Maio)//4
}
