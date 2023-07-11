package str

import (
	"strings"
)

func Concat(input []string) string {
	result := ""
	for _, v := range input {
		result += v + " "
	}

	return result
}

func Marking(input string) string {
	return input[:3] + strings.Repeat("x", len(input[3:]))
}
