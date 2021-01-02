package main

import "fmt"

var z int8 = -128 //-129 will produce error overflows

func main() {
	//numeric
	x := 42
	y := 3.14
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	//string
	s := "wow"
	ss := `That is "Awesome"`
	fmt.Println(s)
	fmt.Println(ss)

	//string is slice of byte
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	fmt.Println()
	for i, v := range s {
		fmt.Printf("at index %d and hexa is %x\n", i, v)
	}
}
