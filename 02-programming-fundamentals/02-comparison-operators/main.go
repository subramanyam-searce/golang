// Using the following operators, write expressions and assign their values to variables:
// g. ==
// h. <=
// i. >=
// j. !=
// k. <
// l. >
// Now print each of the variables.

package main

import (
	"fmt"
)

func main() {
	g := 1 == 3
	h := 3 <= 3
	i := 4 >= 5
	j := 3 != 9
	k := 6 < 10
	l := 5 > 10

	fmt.Println(g, h, i, j, k, l)
}
