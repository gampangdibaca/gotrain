package main

import "fmt"

func main() {
	c := make(<-chan int, 2) //receive only channels

	//will error, does not work
	//c <- 42
	//c <- 43


	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("---------")
	fmt.Printf("%T\n", c)
}
