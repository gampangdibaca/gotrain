package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var mu sync.Mutex
	counter := 0
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println("Counter :\t", counter)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
