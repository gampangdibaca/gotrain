package main

import "fmt"

func main() {
	c := make(chan <- int, 2) //send only channels

	c <- 42
	c <- 43

	//will error, does not work
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	fmt.Println("---------")
	fmt.Printf("%T\n", c)
}
