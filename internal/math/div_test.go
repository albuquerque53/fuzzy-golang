package math

import "testing"

func TestDiv(t *testing.T) {
	type scenery struct {
		n1        int
		n2        int
		expectRes int
	}

	sceneries := []scenery{
		{n1: 10, n2: 2, expectRes: 5},
		{n1: 76, n2: 4, expectRes: 19},
	}

	for _, testScen := range sceneries {
		res := Div(testScen.n1, testScen.n2)

		if res != testScen.expectRes {
			t.Fatalf("Expected %v, received %v", testScen.expectRes, res)
		}
	}
}

func FuzzDiv(f *testing.F) {
	type scenery struct {
		n1        int
		n2        int
		expectRes int
	}

	sceneries := []scenery{
		{n1: 10, n2: 2, expectRes: 5},
		{n1: 76, n2: 4, expectRes: 19},
	}

	for _, testScen := range sceneries {
		f.Add(testScen.n1, testScen.n2)
	}

	f.Fuzz(func(t *testing.T, n1, n2 int) {
		_ = Div(n1, n2)
	})
}
