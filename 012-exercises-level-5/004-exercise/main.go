package main

import "fmt"

func main() {
	m := struct {
		name   string
		weight int
		rgb    bool
	}{
		name:   "Corsair Harpoon",
		weight: 85,
		rgb:    true,
	}
	fmt.Println(m)
	fmt.Println(m.name, m.weight, m.rgb)
}
