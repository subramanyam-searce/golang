package main

import (
	"fmt"
	"testing"

	"github.com/subramanyam-searce/golang/13-testing/01-coverage-pt1/dog"
)

func ExampleYearsTwo() {
	fmt.Println(dog.YearsTwo(21))
	// Output:
	// 147
}

// go test -bench .
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(21)
	}
}
