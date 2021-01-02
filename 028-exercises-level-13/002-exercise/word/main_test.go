package word

import (
	"fmt"
	"github.com/gampangdibaca/gotrain/028-exercises-level-13/002-exercise/quote"
	"strings"
	"testing"
)

func TestCount(t *testing.T) {
	x := Count(quote.SunAlso)
	if x != 7251 {
		t.Error("got", x, "expected", 7251)
	}
}

func TestUseCount(t *testing.T) {
	mapCount := UseCount(quote.SunAlso)
	xs := strings.Fields(quote.SunAlso)
	for k, v := range mapCount {
		wordCount := 0
		for _, v2 := range xs{
			if v2 == k {
				wordCount++
			}
		}
		if wordCount != v {
			t.Error("got", wordCount, "expected", v)
		}
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}

func ExampleCount() {
	fmt.Println(Count(quote.SunAlso))

	// Output:
	// 7251
}
