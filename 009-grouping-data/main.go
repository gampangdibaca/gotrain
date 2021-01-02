package main

import "fmt"

func main() {
	//array
	//array is not idiomatic
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
	//slices
	y := []int{4, 2, 6, 2, 11}
	fmt.Println(y)
	//slice used to group together values of the same type
	//loop in slice with range
	for i, s := range y {
		fmt.Println(i, "is", s)
	}
	//slicing slices
	fmt.Println(y[1:])
	fmt.Println(y[1:3])
	for i := 0; i < len(y); i++ {
		fmt.Println(i, "is", y[i])
	}
	for i, v := range y {
		fmt.Println(i, "is", v)
	}
	//append to slice
	y = append(y, 33, 21, 542, 231)
	fmt.Println(y)

	z := []int{123, 234, 345, 456, 567}
	y = append(y, z...)
	fmt.Println(y)
	//deleting from a slice
	y = append(y[:2], y[4:]...) //works by take the first two elements of slice and append element at index 4 to the end and append it to the first two elements that already taken first
	fmt.Println(y)
	//slices make|| basically make is more efficient, because normal slice will get replaced with a new one when appended. but with make, as long as the capacity not reached, it will keep using same slices
	a := make([]int, 10, 12)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 12)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 13)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 14)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a)) //when cap reached and get appended, the cap doubled
	//multi dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "Vanilla", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut", "Bubble Gum"}
	fmt.Println(mp)
	xp := [][]string{jb, mp}
	fmt.Println(xp)
	//Map
	m := map[string]int{
		"James":      32,
		"Moneypenny": 25,
		"Q":          56,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Q"]) // if no data found will get zero value

	//to check without if
	v, ok := m["Q"]
	fmt.Println(v)
	fmt.Println(ok)

	//to check with if
	if v, ok := m["Q"]; ok {
		fmt.Println("Q is Found and", v, "old")
	}

	//map add & range
	m["bimo"] = 21

	for k, v := range m {
		fmt.Println(k, v)
	}

	//map delete
	delete(m, "bimo")
	fmt.Println(m)
	// deleting not existing key does not produce error
}
