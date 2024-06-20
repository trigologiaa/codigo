package recursividad

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPalindromos(t *testing.T) {
	assert.True(t, EsPalindromo("NEUQUEN"))
	assert.True(t, EsPalindromo("DABALEARROZALAZORRAELABAD"))
	assert.True(t, EsPalindromo("reconocer"))
}

func TestNoPalindromos(t *testing.T) {
	assert.False(t, EsPalindromo("Holanda"))
	assert.False(t, EsPalindromo("estudiar"))
	assert.False(t, EsPalindromo("ANITA LAVA LA TINA"))
	// cómo podemos modificar el código para que reconozca este último palíndromo?
}