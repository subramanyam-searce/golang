package mymath

import (
	"fmt"
	"strconv"
	"testing"
)

type test struct {
	arr []int
	avg float64
}

func TestCenteredAvg(t *testing.T) {
	tests := []test{
		{arr: []int{0, 22, 23, 24, 25, 10000}, avg: 23.5},
		{arr: []int{0, 1, 2, 3, 5, 20}, avg: 2.75},
		{arr: []int{100, 22, 23, 24, 25, 10000}, avg: 43},
		{arr: []int{100, 99, 96, 1, 101}, avg: 98.33},
	}

	for _, v := range tests {
		thisAvg := (CenteredAvg(v.arr))
		thisAvg, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", thisAvg), 64)
		if (thisAvg) != v.avg {
			t.Error("Expected", v.avg, "Got:", thisAvg)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{0, 22, 23, 24, 25, 10000}))
	// Output:
	// 23.5
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{0, 22, 23, 24, 25, 10000})
	}
}
