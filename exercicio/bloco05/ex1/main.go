package main

import (
	"fmt"
	"strconv"
)

func main() {
	notas := []float64{}
	total := float64(0)

	for range 4 {
		var inputTxt string

		fmt.Println("entro com a nota")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		nota, err := strconv.ParseFloat(inputTxt, 64)

		if err != nil {
			fmt.Println("entre com uma nota")
			continue
		}

		notas = append(notas, nota)
		total += nota

	}

	if len(notas) == 0 {
		fmt.Println("nenhuma nota informada")
		return
	}

	media := total / float64(len(notas))

	minimo := notas[0]
	maximo := notas[0]

	for i := range notas {
		if notas[i] < minimo {
			minimo = notas[i]
		}
	}

	for i := range notas {
		if notas[i] > maximo {
			maximo = notas[i]
		}
	}

	fmt.Printf("media: %f, minimo: %f, maximo:%f", media, minimo, maximo)
}

//Faça um programa que receba 4 notas,
//calcule a média, mínimo e máximo dessas notas.
