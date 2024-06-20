package recursividad

func EsPalindromo(texto string) bool {
	if len(texto) <= 1 {
		return true
	}
	primero := texto[0]
	ultimo := texto[len(texto)-1]
	interior := texto[1 : len(texto)-1]
	return primero == ultimo && EsPalindromo(interior)
}