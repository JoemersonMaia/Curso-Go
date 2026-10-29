package main

import "fmt"

func main() {
	var opcao int
	fmt.Println("1-agua com gas\n2-agua sem gas")
	fmt.Scanf("%d", &opcao)
	var qtde float32
	fmt.Println("gostaria de quantas aguas?")
	fmt.Scanf("%f", &qtde)

	switch opcao {
	case 1:
		var valor float32
		fmt.Scanf("%f", &valor)
		valor = 1.5
		fmt.Println("ficou", qtde*valor)
	case 2:
		var valor float32
		fmt.Scanf("%f", &valor)
		valor = 2.5
		fmt.Println("ficou", qtde*valor)
	}
}
