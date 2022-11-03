package math

import "testing"

func TestSum(t *testing.T) {
	type scenery struct {
		n1        int
		n2        int
		expectRes int
	}

	sceneries := []scenery{
		{n1: 2, n2: 3, expectRes: 5},
		{n1: 76, n2: 4, expectRes: 80},
		{n1: 0, n2: 4, expectRes: 4},
	}

	for _, testScen := range sceneries {
		res := Sum(testScen.n1, testScen.n2)

		if res != testScen.expectRes {
			t.Fatalf("Expected %v, received %v", testScen.expectRes, res)
		}
	}
}
