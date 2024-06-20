package division_y_conquista

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPotencias(t *testing.T) {
	assert.Equal(t, 4, Potencia(2, 2))
	assert.Equal(t, 65536, Potencia(2, 16))
	assert.Equal(t, 16777216, Potencia(16, 6))
}