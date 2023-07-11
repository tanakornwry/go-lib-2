package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	input := []int{2, 3, 4}
	result := Sum(input)

	assert.Equal(t, 9, result)
}
