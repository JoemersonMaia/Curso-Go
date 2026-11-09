package main

import "fmt"

func main() {
	lanches := map[string]float64{
		"xburguer": 12.00,
		"xsalada":  15.00,
		"xbacon":   18.00,
	}

	bebidas := map[string]float64{
		"coca":      4.00,
		"fanta":     5.00,
		"antartica": 7.00,
	}

	tamanho_bebida := map[string]float64{
		"pequena": 1.0,
		"media":   1.5,
		"grande":  2.0,
	}

	var lanche string
	fmt.Println("qual o lanche")
	fmt.Scanln(&lanche)

	precolanche, ok := lanches[lanche]
	if !ok {
		fmt.Println("lanche errado")
		return
	}

	var bebida string
	fmt.Println("qual a bebida")
	fmt.Scanln(&bebida)

	precobebida, ok := bebidas[bebida]
	if !ok {
		fmt.Println("bebida errada")
		return
	}

	var tamanho string
	fmt.Println("qual o tamanho da bebida")
	fmt.Scanln(&tamanho)

	precotamanho, ok := tamanho_bebida[tamanho]
	if !ok {
		fmt.Println("tamanho errado")
		return
	}

	total := precolanche + precobebida*precotamanho

	fmt.Println(total)
}
