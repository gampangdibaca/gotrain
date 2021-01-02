package main

import "fmt"

func main() {
	i := 1999
	for {
		fmt.Println(i)
		i++
		if i > 2020 {
			break
		}
	}
}
