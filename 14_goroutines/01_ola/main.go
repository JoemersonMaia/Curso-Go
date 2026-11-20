package main

import (
	"fmt"
	"log"
	"time"
)

func dizAlgo(txt string) {
	for {
		log.Println(txt)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	for i := 1; i <= 100000; i++ {
		go dizAlgo(fmt.Sprintf("%d", i))
	}
	go dizAlgo("ola")
	go dizAlgo("tchau")

	var ok string
	fmt.Scanf("%d", ok)
	log.Println("terminou")
}
