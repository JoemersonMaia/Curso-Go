package main

import "fmt"

func main() {
	var idade int
	fmt.Println("informe sua idade")
	fmt.Scanf("%d", &idade)

	if idade >= 66 {
		fmt.Println("ta bem")
	} else if idade >= 18 {
		fmt.Println("Perfeito")
	} else {
		fmt.Println("cresca")
	}

}
