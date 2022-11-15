package main

import (
	"fmt"

	"github.com/subramanyam-searce/golang/13-testing/02-coverage-pt2/quote"
	"github.com/subramanyam-searce/golang/13-testing/02-coverage-pt2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
