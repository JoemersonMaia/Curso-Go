package main

import (
	"fmt"
	"strconv"
)

func media(total float64, numeros []float64) float64 {
	media := float64(0)
	media = total / float64(len(numeros))
	return media
}

func main() {
	numeros := []float64{}
	total := float64(0)

	for {
		var input string
		fmt.Println("informe um numero")
		fmt.Scanf("%s", &input)

		if input == "" {
			break
		}

		numero, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("entre com um valor valido")
		}
		numeros = append(numeros, numero)
	}
	for _, i := range numeros {
		total += float64(i)
	}
	fmt.Println(media(total, numeros))
}
