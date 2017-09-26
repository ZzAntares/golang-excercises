package main

import "fmt"

func reverse(seq []int) []int {
	if len(seq) <= 1 {
		return seq
	}

	var first int = seq[0]
	var tail []int = seq[1:len(seq)]

	return append(reverse(tail), first)
}

func main() {
	var nums []int = []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(reverse(nums))
}
