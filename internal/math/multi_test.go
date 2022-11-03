package math

import "testing"

func TestMultiply(t *testing.T) {
	type scenery struct {
		n1        int
		n2        int
		expectRes int
	}

	sceneries := []scenery{
		{n1: 3, n2: 2, expectRes: 6},
		{n1: 76, n2: 4, expectRes: 304},
		{n1: 0, n2: 4, expectRes: 0},
	}

	for _, testScen := range sceneries {
		res := Multiply(testScen.n1, testScen.n2)

		if res != testScen.expectRes {
			t.Fatalf("Expected %v, received %v", testScen.expectRes, res)
		}
	}
}
