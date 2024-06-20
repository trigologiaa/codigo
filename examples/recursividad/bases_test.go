package recursividad

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestImprimirEnDecimal(t *testing.T) {
	assert.Equal(t, "127", ImprimirEnDecimal(127))
}

func TestImprimirEnBase(t *testing.T) {
	assert.Equal(t, "101010", ImprimirEnBase(42, 2))
}