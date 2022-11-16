package dog

import (
	"fmt"
	"testing"
)

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(21))
	// Output:
	// 147
}

// go test -bench .
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(21)
	}
}

func ExampleYears() {
	fmt.Println(Years(21))
	// Output:
	// 147
}

// go test -bench .
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(21)
	}
}
