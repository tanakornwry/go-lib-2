package math

func Sum(in []int) int {
	result := 0
	for _, v := range in {
		result += v
	}

	return result
}
