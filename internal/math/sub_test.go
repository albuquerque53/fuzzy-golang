package math

import "testing"

func TestSub(t *testing.T) {
	type scenery struct {
		n1        int
		n2        int
		expectRes int
	}

	sceneries := []scenery{
		{n1: 3, n2: 2, expectRes: 1},
		{n1: 76, n2: 4, expectRes: 72},
		{n1: 0, n2: 4, expectRes: -4},
	}

	for _, testScen := range sceneries {
		res := Sub(testScen.n1, testScen.n2)

		if res != testScen.expectRes {
			t.Fatalf("Expected %v, received %v", testScen.expectRes, res)
		}
	}
}
