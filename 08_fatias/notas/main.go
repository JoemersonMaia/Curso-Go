package main

import (
	"fmt"
	"strconv"
)

func main() {
	notas := []float64{}

	for {
		var inputTxt string
		fmt.Println("entre com a nota")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		nota, err := strconv.ParseFloat(inputTxt, 64)
		if err != nil {
			fmt.Println("entre com uma nota valida")
		}
		notas = append(notas, nota)

	}
	fmt.Println(notas)
}
