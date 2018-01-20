package main

import (
	"fmt"
	"sync"
)

var count int
var lock sync.Mutex

func increment()  {
	lock.Lock()
	defer lock.Unlock()
	count++
	fmt.Println("incremented:", count)
}

func main()  {
	wg := sync.WaitGroup{}

	for i := 0; i< 10; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("finished:", count)
}