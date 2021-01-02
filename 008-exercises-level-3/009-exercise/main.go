package main

import "fmt"

func main() {
	favSport := "Badminton"
	switch favSport {
	case "Swimming":
		fmt.Println("Go To Pool")
	case "Badminton":
		fmt.Println("Go To Hall")
	case "Soccer":
		fmt.Println("Go To Field")
	}
}
