package base 

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func M_closure() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // Output: 1
	fmt.Println(nextInt()) // Output: 2
	fmt.Println(nextInt()) // Output: 3

	nextInts := intSeq()
	fmt.Println(nextInts()) // Output: 1
	fmt.Println(nextInts()) // Output: 2
	fmt.Println(nextInts()) // Output: 3
}
