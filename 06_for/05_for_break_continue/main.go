package main

import "fmt"

func main() {
	nome := "teo calvo"

	for _, v := range nome {
		letra := (string(v))
		if letra == "l" {
			break
		}

		if letra == " " {
			continue

		}
		fmt.Println(letra)
	}

}
