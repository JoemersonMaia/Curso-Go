package main

import (
	"fmt"
	"slices"
)

func main() {
	notas := []float64{}

	for i := 1; i <= 4; i++ {
		var nota float64
		fmt.Println("entre com a nota")
		fmt.Scanf("%f", &nota)
		notas[i-1] = nota
	}

	min := slices.Min(notas[:])
	max := slices.Max(notas[:])

	var total float64
	for _, v := range notas {
		total += v
	}

	fmt.Println("mdia:", total/float64(len(notas)))
	fmt.Println("minimo:", min)
	fmt.Println("minimo:", max)
}

//Faça um programa que receba 4 notas,
//calcule a média, mínimo e máximo dessas notas.
