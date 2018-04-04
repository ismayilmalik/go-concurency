package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	close(ch)

	go sendToChan(ch)

	val, ok := <-ch
	fmt.Println(val, ok)
}

func sendToChan(chnl chan int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	chnl <- 0 //panic: send on closed channel
}
