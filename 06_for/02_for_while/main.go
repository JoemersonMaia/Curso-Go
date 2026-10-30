package main

import "fmt"

func main() {
	for i := 1; i <= 1000; i++ {
		if i%25 == 0 {
			fmt.Println(i)
		}
	}
}
