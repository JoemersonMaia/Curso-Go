package main

import "fmt"

func primeira() {
	fmt.Println("eu sou a primeira")
}
func segunda() {
	fmt.Println("eu sou a segunda")
}
func ultima() {
	fmt.Println("eu sou a ultima")
}
func main() {
	defer ultima()
	defer primeira()
	segunda()
}
