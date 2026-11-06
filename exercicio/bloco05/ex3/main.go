package main

import "fmt"

func main() {
	vendas := [6]float64{}
	total := float64(0)

	for i := 1; i <= 6; i++ {
		var venda float64
		fmt.Printf("informe o valor de venda do %d", i)
		fmt.Scanf("%f", &venda)
		vendas[i-1] = venda
	}

	newvendas := vendas[1 : len(vendas)-1]

	for _, i := range newvendas {
		total += float64(i)
	}

	media := total / float64(len(newvendas))
	fmt.Println("media=", media)

	if media < 1000 {
		fmt.Println("temos que vender mais, esta pouco")
	} else {
		fmt.Println("e isso tao bem dms")
	}
}
