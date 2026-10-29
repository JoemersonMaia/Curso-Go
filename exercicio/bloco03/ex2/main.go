package main

import "fmt"

func main() {
	var opcao int
	fmt.Println("1-casquinha (R$1,00)\n2-cascão (R$2,50)\n3-cestinha (R$4,00)")
	fmt.Scanf("%d", &opcao)
	var opcaocopo float32

	switch opcaocopo {
	case 1:
		fmt.Scanf("%f", &opcaocopo)
		opcaocopo = 1.0
	case 2:
		fmt.Scanf("%f", &opcaocopo)
		opcaocopo = 2.5
	case 3:
		fmt.Scanf("%f", &opcaocopo)
		opcaocopo = 4.0
	}

	var opcaosabor int
	fmt.Println(" 1-morango\n2-creme\n3-chocolate")
	fmt.Scanf("%d", &opcaosabor)
	switch opcaosabor {
	case 1:
		var sabor string
		fmt.Scanf("%s", &sabor)
		sabor = "morango"
	case 2:
		var sabor string
		fmt.Scanf("%s", &sabor)
		sabor = "morango"
	case 3:
		var sabor string
		fmt.Scanf("%s", &sabor)
		sabor = "chocolate"
	}

	var opcaocorbetura int
	fmt.Println("1-Caramelo (R$1,50)\n2-morango (R$1,50)\n3-chocolate (R$1,50)\n4-sem cobertura (R$0,00)")
	fmt.Scanf("%d", &opcaocorbetura)
	var corbetura float32

	switch opcaocorbetura {
	case 1:
		fmt.Scanf("%s", &corbetura)
		corbetura = 1.5
	case 2:
		fmt.Scanf("%s", &corbetura)
		corbetura = 1.5
	case 3:
		fmt.Scanf("%s", &corbetura)
		corbetura = 1.5
	case 4:
		fmt.Scanf("%f", &corbetura)
		corbetura = 0.0
	}
	var valor float32
	fmt.Scanf("%f", &valor)
	valor = opcaocopo + corbetura
	fmt.Printf("o valor a ser pago é %f", valor)

}
