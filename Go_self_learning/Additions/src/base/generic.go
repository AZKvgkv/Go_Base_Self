package base

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func M_generic() {
	var m = map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println("keys m:", MapKeys(m))

	_=MapKeys[int, string](m)

	lst:=List[int]{}
	lst.Push(10)
	lst.Push(20)
	lst.Push(30)
	fmt.Println("list:", lst.GetAll())
}
