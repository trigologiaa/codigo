package recursividad

// Factorial calcula el factorial de n, donde n debe ser mayor o igual a 0.
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}