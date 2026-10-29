package main

import "fmt"

func main() {
	var idade int
	fmt.Println("informe sua idade")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		println("PERFEITO")
	}

	if idade <= 17 {
		fmt.Println("cresca")
	}
}
