package tester

import "testing"

var flagtests = []struct {
	in  []int
	out int
}{
	{[]int{1, 3, 5, 7}, 16},
	{[]int{0, 0, 0, 0}, 0},
	{[]int{-5, 0, -5, 3}, -7},
}

func TestSum(t *testing.T) {
	for _, testCase := range flagtests {
		var sum int = Sum(testCase.in)
		if sum != testCase.out {
			t.Logf(
				"Expected sum of %v to be %d but was %d",
				testCase.in,
				testCase.out, sum)
			t.Fail()
		}
	}
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 3, 5, 7, 9})
	}
}
