package main

import (
	"fmt"
	"sync"
)

var count int = 0

func increment(){
	count++
}

func main()  {
	once := sync.Once{}
	wg := sync.WaitGroup{}
	wg.Add(100)

	for index := 0; index < 100; index++ {
		go func(){
			defer wg.Done()
			once.Do(increment)
		}()
	}

	wg.Wait()
	fmt.Println("Count:", count)
}