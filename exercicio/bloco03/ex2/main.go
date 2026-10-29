package main

import "fmt"

func main() {
	var tiposorvete, saborsorvete, corbetura string
	fmt.Println("casquinha (R$1,00)\ncascão (R$2,50)\ncestinha (R$4,00)")
	fmt.Scanf("%s", &tiposorvete)

	fmt.Println("caramelo (R$1,50)\nMorango (R$1,50)\nChocolate (R$1,50)\n4-Sem cobertura (R$0,00)")
	fmt.Scanf("%s", &saborsorvete)

	fmt.Println("Qual cobertura? [caramelo/morango/chocolate/sem cobertura]")
	fmt.Scanf("%s", &corbetura)

	var valortipo, valorsabor float64

	switch tiposorvete {
	case "casquinha":
		valortipo = 1.0
	case "cascao":
		valortipo = 2.5
	case "cestinha":
		valortipo = 4.0
	default:
		fmt.Println("coloque um valro valido")
	}

	if corbetura == "caramelo" || corbetura == "morango" || corbetura == "chocolate" {
		valorsabor = 1.5
	} else if corbetura == "sem corbetura" {
		valorsabor = 0
	} else {
		fmt.Println("coloque um valor valido")
	}

	total := valorsabor + valortipo
	fmt.Printf("total %2.f", total)
}
