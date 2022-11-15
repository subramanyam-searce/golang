package word

import (
	"strings"

	"github.com/subramanyam-searce/golang/13-testing/02-coverage-pt2/quote"
)

// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

func Count(s string) int {
	m := UseCount(quote.SunAlso)

	return m[s]
}
