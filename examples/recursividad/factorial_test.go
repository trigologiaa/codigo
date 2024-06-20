package recursividad

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFactorialDeCero(t *testing.T) {
	assert.Equal(t, 1, Factorial(0))
}

func TestFactorialDe5(t *testing.T) {
	assert.Equal(t, 120, Factorial(5))
}