package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("me informe um numeoro")
	var n int
	fmt.Scanf("%d", &n)
	resultado := math.Sqrt(float64(n))
	fmt.Printf("ao quadrado é: %.2f\n", resultado)
}
