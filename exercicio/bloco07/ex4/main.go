package main

import (
	"fmt"
	"math"
)

func quadrado(numero float64) float64 {
	quadrado := math.Pow(numero, 2)
	return quadrado
}

func main() {
	var numero float64
	fmt.Println("informe um numero")
	fmt.Scanf("%f", &numero)

	fmt.Println(quadrado(numero))
}
