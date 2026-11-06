package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	total := float64(0)
	numeros := []float64{}

	for {
		var inputTxt string
		fmt.Println("informe um valor")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		numero, err := strconv.ParseFloat(inputTxt, 64)
		if err != nil {
			fmt.Println("entre com um valor valido")
		}

		numeros = append(numeros, numero)

	}

	for _, i := range numeros {
		total += float64(i)
	}
	min := slices.Min(numeros[:])
	max := slices.Max(numeros[:])

	fmt.Println(total)
	fmt.Println(min)
	fmt.Println(max)
}
