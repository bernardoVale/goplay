package main

import "fmt"

func listAnything( data ...interface{}){

	fmt.Println("Listando qualquer coisa")
	for _, d := range data{

		fmt.Println("Value:", d)
	}
}
func listStrings(data ...string){
	for _, d := range data{
		fmt.Println("Value:", d)
	}

}
func main() {
	mylist := []string{"capa", "de", "pica"}


	// É tipo o kwargs do python
	listStrings("capa", "de", "pica")

	//mylist é um slice mas eu posso envia-lo ao listAnything
	listAnything(mylist,[]string{"capudo", "animal"}, []int{2,3}, []float64{4.4, 6.7, 4.3})
}
