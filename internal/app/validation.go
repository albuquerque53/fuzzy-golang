package app

import (
	"fuzzygolang/internal/util"
	"log"
	"strconv"
)

func validateOperation(op string) string {
	validOps := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	if !validOps[op] {
		log.Fatal("Invalid operation. Possible values are +, -, * and /")
	}

	return op
}

func validateNumbers(n1, n2 string) (int, int) {
	intN1, err := strconv.Atoi(n1)

	util.HandleError(err)

	intN2, err := strconv.Atoi(n2)

	util.HandleError(err)

	return intN1, intN2
}
