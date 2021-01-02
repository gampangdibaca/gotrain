package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	sa := [][]string{
		jb,
		mp,
	}
	for _, v := range sa {
		fmt.Println(v)
		for _, va := range v {
			fmt.Println(va)
		}
	}
}
