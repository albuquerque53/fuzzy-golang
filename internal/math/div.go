package math

// Div must return the division of n1 by n2
// if n2 equals to 0, 0 is returned
func Div(n1, n2 int) int {
	if n2 == 0 {
		return 0
	}

	return n1 / n2
}
