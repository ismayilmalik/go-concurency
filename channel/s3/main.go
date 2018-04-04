package main

import (
	"fmt"
)

func main()  {
	ch := buffChanOwner()
	for v := range ch {
		fmt.Printf("Received: %d\n", v)
	}
	fmt.Println("Done!")
}

func buffChanOwner() <-chan int {
	chancap := 4
	ch := make(chan int, chancap)
	go func ()  {
		defer close(ch) //if we comment this line for loop in main func will panic
		for i := 0; i < chancap * 2; i++ {
			fmt.Printf("Sending: %d\n", i)
			ch <- i
		}
	}()
	return ch
}