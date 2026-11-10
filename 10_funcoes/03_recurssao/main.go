package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}
func main() {
	fmt.Println("fib(o)", fib(0))
	fmt.Println("fib(1)", fib(1))
	fmt.Println("fib(2)", fib(2))
	fmt.Println("fib(3)", fib(3))
	fmt.Println("fib(4)", fib(4))
}
