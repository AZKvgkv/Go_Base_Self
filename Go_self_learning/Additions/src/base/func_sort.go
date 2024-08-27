package base

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func M_func_sort() {
	fruits := []string{"apple", "banana", "orange", "kiwi", "grape"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}