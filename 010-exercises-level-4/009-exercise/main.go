package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"Moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	m["bimo"] = []string{`Pizza`, `Code`, `Tea`}
	for k, v := range m {
		fmt.Println(k)
		for i, v2 := range v {
			fmt.Printf("\t%d, %s\n", i, v2)
		}
	}
}
