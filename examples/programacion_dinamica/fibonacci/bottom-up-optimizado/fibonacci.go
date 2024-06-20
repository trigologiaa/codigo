package fibbonacci_bottom_up_optimizado

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	prev2 := 0
	prev := 1
	var current int
	for i := 2; i <= n; i++ {
		//current, prev2, prev = prev2+prev, prev, current
		current = prev2 + prev
		prev2 = prev
		prev = current
	}
	return current
}