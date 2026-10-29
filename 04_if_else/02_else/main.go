package main

import "fmt"

func main() {
	var idade int
	fmt.Println("informe sua idade")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		println("PERFEITO")
	} else {
		fmt.Println("cresca")
	}

}
