package fibbonacci_top_down

func Fibonacci(n int) int {
	mem := make(map[int]int)
	mem[0] = 0
	mem[1] = 1
	return fibonacci(n, mem)
}

func fibonacci(n int, mem map[int]int) int {
	key := n
	result, ok := mem[key]
	if ok {
		return result
	}
	mem[key] = fibonacci(n-1, mem) + fibonacci(n-2, mem)
	return mem[key]
}