package main

import (
	"fmt"
	"runtime"
)

var counter int
var counter2 int

func main() {
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go send(c)
	}
	fmt.Println("Num of Goroutines:", runtime.NumGoroutine())
	receive(c)
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
		counter2++
	}
	counter++
	if counter == 10 {
		close(c)
	}
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(counter2)
}
