package main

import (
	"sync"
	"time"
	"fmt"
)

func main()  {
	wg:= sync.WaitGroup{}
	wg.Add(1)

	asyncLog("some message", func(msg string){
		defer wg.Done()
		fmt.Println(msg)
	})
}

func asyncLog(msg string, appenderCb func(string))  {
	logMsg := fmt.Sprintf("Time: %s, Context: %s", time.Now(), msg)
	go func(){
		appenderCb(logMsg)
	}()
}