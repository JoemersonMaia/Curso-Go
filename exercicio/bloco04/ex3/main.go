package main

import "fmt"

func main() {
	total := 0.0
	for {
		var saldo float64
		fmt.Println("informe o valor que gsotaria de adicionar")
		fmt.Scanf("%f", &saldo)

		total = total + saldo
		if saldo == 0 {
			fmt.Printf("total adicioando a conta: %f", total)
			break
		}
	}
}
