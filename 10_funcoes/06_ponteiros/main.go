package main

import "fmt"

func add0ne(a int) int {
	a++
	return a + 1
}

func add0nePt(a *int) {
	*a++
}
func main() {
	numero := 1
	valor := add0ne(numero)
	add0nePt(&numero)
	fmt.Println(numero)
	fmt.Println(valor)
}
