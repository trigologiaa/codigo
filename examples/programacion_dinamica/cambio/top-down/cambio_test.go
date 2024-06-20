package cambio_top_down

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCambio(t *testing.T) {
	assert.Equal(t, 3, Cambio(15, []int{1, 5, 11}))
}