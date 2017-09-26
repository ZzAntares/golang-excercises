package main

import "fmt"

func chopper(seq []int) func() ([]int, []int) {
	var x int = 1

	return func() (former, last []int) {
		former = seq[0:x]
		last = seq[x:len(seq)]
		x++
		return
	}
}

func add(seq []int) int {
	var s int

	for _, el := range seq {
		s += el
	}

	return s
}

func balancer(seq []int) bool {
	var chop func() ([]int, []int) = chopper(seq)

	for i := 0; i < len(seq)-1; i++ {
		var head, tail []int = chop()

		if add(head) == add(tail) {
			return true
		}
	}

	return false
}

func main() {
	var test1 []int = []int{1, 1, 1, 2, 1}
	fmt.Printf("%v -> %v\n", test1, balancer(test1))

	var test2 []int = []int{10, 10}
	fmt.Printf("%v -> %v\n", test2, balancer(test2))

	var test3 []int = []int{2, 1, 1, 2, 1}
	fmt.Printf("%v -> %v\n", test3, balancer(test3))
}
