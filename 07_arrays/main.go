package main

import "fmt"

func main() {
	var x [20]float64
	fmt.Println(x, "\n")

	var y = [20]float64{1, 22, 32, 43, 53}
	y[18] = 33
	fmt.Println("y =", y)
	fmt.Println("elemento de y = y[0] = ", y[0])
	fmt.Println("elemento de y = y[0:2] = ", y[0:2])
	fmt.Println("elemento de y = y[18:20] = ", y[18:20])
	fmt.Println("elemento de y = y[len(y)-2:] = ", y[len(y)-2:])

	var total float64
	for _, v := range y {
		total += v
	}
	fmt.Println("total", total)

	t := [10]string{"teo", "jojo"}
	fmt.Println("\n", t)

	z := [5]uint64{32, 432, 543, 312, 43}
	fmt.Println(z)

}
