package recursividad

import "fmt"

func ImprimirEnDecimal(n int) string {
	res := ""
	if n >= 10 {
		res += ImprimirEnDecimal(n / 10)
	}
	return res + fmt.Sprintf("%d", n%10)
}

const digits string = "0123456789abcdef"

func ImprimirEnBase(n int, base int) string {
	res := ""
	if n >= base {
		res += ImprimirEnBase(n/base, base)
	}
	return res + string(digits[n%base])
}