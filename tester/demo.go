package tester

func Sum(seq []int) int {
	var total int

	for _, el := range seq {
		total += el
	}

	return total
}
