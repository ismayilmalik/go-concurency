package main

import (
	"fmt"
	"sync"
)

func main()  {
	wg := sync.WaitGroup{}
	wg.Add(1)

	// fork point
	go func ()  {
		defer wg.Done()
		fmt.Println("From child...")
	}()

	fmt.Println("From parent...")

	// join point
	wg.Wait()
}