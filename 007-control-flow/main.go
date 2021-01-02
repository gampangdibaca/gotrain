package main

import "fmt"

func main() {
	//Loop
	//Standard Loop
	for i := 0; i <= 100; i++ {
		fmt.Println("Hello", i)
	}
	//Nesting Loop
	for i := 0; i <= 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(i, j)
		}
	}
	//For statement with single conditions
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")

	//for statement with for clause
	for i := 0; i <= 5; i++ {
		fmt.Println("Hello", i)
	}

	//for statement with range clause
	s := "James_Bond"
	for i, c := range []byte(s) {
		fmt.Printf("Index %d : %T\n", i, c)
	}

	//Break & Continue
	for i := 0; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("Break & Continue Separator")
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//Printing ASCII
	for i := 32; i <= 126; i++ {
		fmt.Printf("%q\n", i)
	}
	//if statement
	if true {
		fmt.Println("Its True")
	}
	if false {
		fmt.Println("Its False")
	}
	if !true {
		fmt.Println("Its True")
	}
	if !false {
		fmt.Println("Its False")
	}
	if 2 == 2 {
		fmt.Println("Its True")
	}
	if 2 != 2 {
		fmt.Println("Its False")
	}
	if !(2 == 2) {
		fmt.Println("Its True")
	}
	if !(2 != 2) {
		fmt.Println("Its False")
	}
	if x := 44; x == 2 { //initialization statement, x only available inside if block
		fmt.Println(x)
	}
	//if , else if, else
	a := 42
	if a == 40 {
		fmt.Println("x ==", a)
	} else if a == 42 {
		fmt.Println("x ==", a)
	} else {
		fmt.Println("x !=", a)
	}
	//Loop conditional modulus
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
	//Switch statement
	switch {
	case false:
		fmt.Println("not print")
	case (2 == 4):
		fmt.Println("not print2")
	case (3 == 3):
		fmt.Println("print")
		fallthrough //fallthrough not the best practice
	case (4 == 4):
		fmt.Println("print2")
		fallthrough
	case (7 == 9):
		fmt.Println("not print3")
		fallthrough
	case (3 == 5):
		fmt.Println("not print4")
		fallthrough
	case (3 == 3):
		fmt.Println("print3")
		fallthrough
	default:
		fmt.Println("Default")
	}
	//switch on value
	// n := "Bond"  <-- also possible
	switch "Bond" {
	case "Kuku", "Bond": //multiple switch value
		fmt.Println("Finger")
	case "Bondd":
		fmt.Println("James Bond")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("Default")
	}
	//conditional logic operators
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}
