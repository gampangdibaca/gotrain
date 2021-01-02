package main

import "fmt"

type customErr struct {
	msg string
	err error
}

func (c customErr) Error() string {
	return fmt.Sprintf("Just an awesome custom error :D")
}

func main() {
	c := customErr{
		msg: "The Custom ERROR",
		err: fmt.Errorf("Just a normal error"),
	}
	foo(c)
}

func foo(err error) {
	fmt.Println(err)
}
