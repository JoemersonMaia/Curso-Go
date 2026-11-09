package main

import "fmt"

func main() {
	sorvetes := map[string]float64{
		"casquinha": 1.0,
		"cascao":    2.5,
		"cestinha":  4.0,
	}

	corbeturas := map[string]float64{
		"caramelo":  1.5,
		"morango":   1.5,
		"chocolate": 1.5,
		"cobertura": 0.0,
	}

	var sorvete string
	fmt.Println("informe a casquinha")
	fmt.Scanf("%s", &sorvete)

	precocasquinha, ok := sorvetes[sorvete]
	if !ok {
		fmt.Println("casquinha invalida")
		return
	}

	var corbetura string
	fmt.Println("informe cobertura")
	fmt.Scanf("%s", &corbetura)

	precocorbetura, ok := corbeturas[corbetura]
	if !ok {
		fmt.Println("corbetura invalida")
		return
	}

	total := precocasquinha + precocorbetura

	fmt.Println(total)
}
