package main

import "fmt"

func main() {
	var opcao int
	fmt.Println("Entre com uma opcao")
	fmt.Scanf("%d", &opcao)

	switch opcao {
	case 1:
		fmt.Println("banana")
	case 2:
		fmt.Println("maca")
	case 3:
		fmt.Println("pera")
	default:
		fmt.Println("uva")
	}
}
