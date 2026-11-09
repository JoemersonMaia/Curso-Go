package main

import "fmt"

func main() {
	idades := make(map[string]int)
	idades["teo"] = 33
	idades["nah"] = 34
	fmt.Println(idades)

	alturas := map[string]float64{}
	alturas["teo"] = 1.82
	fmt.Println(alturas)

	alturaTeo := alturas["teo"]
	fmt.Println(alturaTeo)

	alturaNah, ok := alturas["nah"]
	if ok {
		fmt.Println(alturaNah)
	} else {
		fmt.Println("ta errado")
	}

	if valor, ok := alturas["lara"]; ok {
		fmt.Println("a altura da lara é:", valor)
	} else {
		fmt.Println("nao encontrado")
	}

}
