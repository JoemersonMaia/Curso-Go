package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {
	temps := []float64{}
	total := float64(0)

	for range 6 {
		var inputTxt string
		fmt.Println("escreva uma temp")
		fmt.Scanf("%s", &inputTxt)

		if inputTxt == "" {
			break
		}

		temp, err := strconv.ParseFloat(inputTxt, 64)

		if err != nil {
			fmt.Println("entre com uma nota correta pfvr")
			continue
		}

		temps = append(temps, temp)

	}

	temps = slices.Delete(temps, 0, 1)
	temps = slices.Delete(temps, len(temps)-1, len(temps))

	for _, i := range temps {
		total += float64(i)
	}

	media := total / float64(len(temps))

	if media > 30 {
		fmt.Printf("esta calor\n")
	} else {
		fmt.Println("ta frio\n")
	}
}
