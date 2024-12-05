package calcs

import "testing"

type test struct {
	arg1     []int
	expected int
}

var testsIncreasing = []test{
	{[]int{1, 2, 3, 4, 5}, 1},
	{[]int{1, 2, 3, 5, 5}, 0},
	{[]int{1, 3, 2, 5, 5}, 0},
}

var testsDecreasing = []test{
	{[]int{5, 4, 3, 2, 1}, 1},
	{[]int{5, 4, 3, 7, 1}, 0},
	{[]int{5, 5, 3, 2, 1}, 0},
}

func testIsSafeDecreasing(t *testing.T) {
	for _, test := range testsDecreasing {
		if output := IsSafeDecreasing(test.arg1); output != test.expected {
			t.Errorf("Output %d not equal to expected %d.", output, test.expected)
		}
	}
}
