package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data []int
		answer float64
	}

	tests := []test{
		{[]int{1, 4, 6, 8, 100}, 6,},
		{[]int{0, 8, 10, 1000}, 9,},
		{[]int{9000, 4, 10, 8, 6, 12}, 9,},
		{[]int{123, 744, 140, 200}, 170,},
	}

	for _, v := range tests{
		cavg := CenteredAvg(v.data)
		if cavg != v.answer {
			t.Error("got", cavg, "expected", v.answer)
		}
	}

}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := []int{1, 4, 6, 8, 100}
		CenteredAvg(a)
	}
}

func ExampleCenteredAvg() {
	a := []int{1, 4, 6, 8, 100}
	fmt.Println(CenteredAvg(a))

	// Output:
	// 6
}
