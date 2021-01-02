package main

import (
	"fmt"
	"github.com/gampangdibaca/gotrain/028-exercises-level-13/002-exercise/quote"
	"github.com/gampangdibaca/gotrain/028-exercises-level-13/002-exercise/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
