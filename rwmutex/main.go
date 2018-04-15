package main

import (
	"fmt"
	"sync"
)

type DataStorage struct {
	sync.RWMutex
	store map[string]interface{}
}

func (d *DataStorage) set(key string, val interface{})  {
	d.Lock()
	defer d.Unlock()
	d.store[key] = val
}

func (d *DataStorage) get(key string) interface{} {
	d.RLock()
	defer d.RUnlock()
	return d.store[key]
}

func main()  {
	d := &DataStorage{ store: make(map[string]interface{})}
	wg := sync.WaitGroup{}
	wg.Add(5)
	go func() {
		defer wg.Done()
		d.set("MT", "Mayk Tyson")
	}()
	for index := 0; index < 4; index++ {
		go func(){
			defer wg.Done()
			fmt.Println(d.get("MT"), "read lock")
		}()
	}
	wg.Wait()
}