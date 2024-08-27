package base

import (
	"fmt"
	"sort"
)

func M_sort() {
	str_s := []string{"c", "a", "b"}
	sort.Strings(str_s)
	fmt.Printf("Sorted string slice: %v\n", str_s)

	int_s:=[]int{7,2,4}
	sort.Ints(int_s)
	fmt.Printf("Sorted int slice: %v\n", int_s)

	s:=sort.IntsAreSorted(int_s)
	fmt.Println("Sorted: ", s)
}
