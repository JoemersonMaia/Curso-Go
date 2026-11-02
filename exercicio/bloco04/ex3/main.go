package main

import (
	"fmt"
	"strconv"
)

func main() {
	var soma float64
	for {
		var entrada string
		fmt.Println("informe o valor que gsotaria de adicionar")
		fmt.Scanf("%s", &entrada)

		if entrada == "" {
			break
		}

		valor, erro := strconv.ParseFloat(entrada, 64)
		if erro != nil {
			fmt.Println("ente com um valor avalido")
			continue
		}
		soma += valor
	}
	fmt.Printf("saldo %f", soma)
}
