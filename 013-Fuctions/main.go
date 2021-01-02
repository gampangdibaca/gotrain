package main

import "fmt"

func main() {
	foo()
	bar("Hello")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Bimo", "PK")
	fmt.Println(x)
	fmt.Println(y)
	z := goo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(z)
	xi := []int{22, 3, 4, 5, 6, 7, 8, 9}
	zz := goo(xi...)
	fmt.Println(zz)
	//defer, defer used to run a func when other func already finish. usually used to make sure a func will be run
	defer doo()
	soo()

	//anonymous function
	func() {
		fmt.Println("Anonymous func")
	}()
	func(x int) {
		fmt.Println("Anonymous func", x)
	}(42)

	//func expression
	f := func() {
		fmt.Println("func expression")
	}
	f()
	fe := func(x int) {
		fmt.Println("func expression says", x)
	}
	fe(43)
	//returning func from func
	m1 := joo()
	x1 := m1()
	fmt.Println(x1)
	//callback is passing func as an argument
	ii := []int{1, 2, 3, 4, 5, 6, 222}
	s := sum(ii...)
	fmt.Println(s)
	s2 := even(sum, ii...)
	fmt.Println("even number only is", s2)
	s3 := odd(sum, ii...)
	fmt.Println("odd number only is", s3)
	//closure
	{
		var y2 int = 42
		fmt.Println(y2)
	}
	//fmt.Println(y2) <--- will error , out of scope
	//recursion
	fmt.Println(factorial(5))
	fmt.Println(forfacto(5))
}

// func (r receiver) identifier (parameters) (return(s)) {...}
func foo() {
	fmt.Println("From foo")
}

//  Go is PASS BY VALUE
func bar(s string) {
	fmt.Println("Passed message is", s)
}

//returning value from func
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, ln, `, says "Hello"`)
	b := false
	return a, b
}

//variadic parameters
func goo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
	return sum
}

//defer
func doo() {
	fmt.Println("doo")
}

func soo() {
	fmt.Println("soo")
}

//returning func from func
func joo() func() int {
	return func() int {
		return 444
	}
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

//callback
func even(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
func odd(f func(xi ...int) int, vi ...int) int {
	yi := []int{}
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func forfacto(n int) int {
	x := 1
	for ; n > 0; n-- {
		x *= n
	}
	return x
}
