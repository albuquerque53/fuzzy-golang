package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func AskValue(question string) string {
	var value string

	input := bufio.NewReader(os.Stdin)

	for {
		fmt.Fprint(os.Stderr, question)
		value, _ = input.ReadString('\n')

		if value != "" {
			break
		}
	}

	return strings.TrimSpace(value)
}
