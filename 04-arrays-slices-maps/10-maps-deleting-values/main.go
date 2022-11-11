// Using the code from the previous example, delete a record from your map. Now print the map
// out using the “range” loop

package main

import "fmt"

func main() {
	mp := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(mp, "no_dr")

	for k, v := range mp {
		fmt.Println(k, v)
	}
}
