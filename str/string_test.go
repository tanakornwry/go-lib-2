package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConcat(t *testing.T) {
	input := []string{"Hello", "world"}
	result := Concat(input)

	assert.Equal(t, "Hello world ", result)
}

func TestMarking(t *testing.T) {
	result := Marking("Tanakorn")
	assert.Equal(t, "Tanxxxxx", result)
}
