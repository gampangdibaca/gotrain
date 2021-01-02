package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i < 3 {
			fmt.Println("small")
		} else if i < 8 {
			fmt.Println("medium")
		} else {
			fmt.Println("big")
		}

	}
}
