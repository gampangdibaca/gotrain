package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(3)
	if x != 21 {
		t.Error("got", x, "expected", 21)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(3)
	if x != 21 {
		t.Error("got", x, "expected", 21)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))

	// Output:
	// 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))

	// Output:
	// 21
}