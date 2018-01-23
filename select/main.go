package main

import (
	"fmt"
	"time"
)

func sum(c, q chan int) {
	var sum int

	for {
		select {
		case x := <-c:
			sum += x
		case <-q:
			fmt.Println("Total:", sum)
			return
		}
	}
}

func main() {
	ch := make(chan int)
	qu := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 10)
			ch <- i
		}
		qu <- 0
	}()

	sum(ch, qu)
}
