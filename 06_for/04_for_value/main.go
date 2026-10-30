package main

import "fmt"

func main() {
	nome := "teo calvo"

	for _, v := range nome {
		fmt.Println(string(v))
	}

}
