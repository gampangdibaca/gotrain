package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%d\n", i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%U %q\n", i, i)
		}
	}
}
