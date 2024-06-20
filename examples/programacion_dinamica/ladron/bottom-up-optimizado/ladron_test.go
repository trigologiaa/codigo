package bottom_up_optimizado

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLadron(t *testing.T) {
	assert.Equal(t, 25, Rob([]int{2, 10, 3, 6, 8, 1, 7}))
}