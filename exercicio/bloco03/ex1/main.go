package main

import "fmt"

func main() {
	var opcao int
	fmt.Println("1-agua com gas\n2-agua sem gas")
	fmt.Scanf("%d", &opcao)
	var qtde float64
	fmt.Println("gostaria de quantas aguas?")
	fmt.Scanf("%f", &qtde)

	switch opcao {
	case 1:
		fmt.Println(1.5 * qtde)
	case 2:
		fmt.Println(2.5 * qtde)
	default:
		fmt.Println("entrte com 1 ou 2")
	}
}
