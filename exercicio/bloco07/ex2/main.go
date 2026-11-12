package main

import "fmt"

func ehpar(numero int64) int64 {
	if numero%2 == 0 {
		fmt.Println("ehpar")
	} else {
		fmt.Println("ehimpar")
	}
	return numero
}

func main() {
	var numero int64
	fmt.Println("informe um numero")
	fmt.Scanf("%d", &numero)
	ehpar(numero)
}
