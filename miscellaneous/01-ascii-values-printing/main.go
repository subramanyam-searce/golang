package main

import "fmt"

func divider() {
	fmt.Println("------------------------")
}

func main() {
	//int to uppercase charcters
	for i := 65; i < 91; i++ {
		fmt.Printf("%c\n", i)
	}
	divider()

	//int to lowercase charecters
	for i := 97; i < 123; i++ {
		fmt.Printf("%c\n", i)
	}
	divider()

	//vice versa
	fmt.Printf("%d\n", 'A')

	//string shall be converted to slice of byte
	fmt.Println([]byte("AB"))
}
