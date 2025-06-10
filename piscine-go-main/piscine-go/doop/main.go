package main

import "os"

const (
	MaxInt = 9223372036854775807
	MinInt = -9223372036854775808
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func atoi(s string) (int64, bool) {
	var n int64
	sign := int64(1)
	i := 0

	if len(s) == 0 {
		return 0, false
	}

	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	for ; i < len(s); i++ {
		if !isDigit(s[i]) {
			return 0, false
		}
		d := int64(s[i] - '0')
		if n > (MaxInt-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}

	n *= sign
	if n < MinInt || n > MaxInt {
		return 0, false
	}
	return n, true
}

func printInt64(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte("0\n"))
		return
	}
	var buf [20]byte
	i := len(buf) - 1
	neg := n < 0
	if neg {
		n = -n
	}
	for n > 0 {
		buf[i] = byte('0' + n%10)
		n /= 10
		i--
	}
	if neg {
		buf[i] = '-'
		i--
	}
	os.Stdout.Write(buf[i+1:])
	os.Stdout.Write([]byte("\n"))
}

func printString(s string) {
	os.Stdout.Write([]byte(s + "\n"))
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok1 := atoi(aStr)
	b, ok2 := atoi(bStr)

	if !ok1 || !ok2 {
		return
	}

	switch op {
	case "+":
		if (b > 0 && a > MaxInt-b) || (b < 0 && a < MinInt-b) {
			return
		}
		printInt64(a + b)
	case "-":
		if (b < 0 && a > MaxInt+b) || (b > 0 && a < MinInt+b) {
			return
		}
		printInt64(a - b)
	case "*":
		if (a > 0 && b > 0 && a > MaxInt/b) ||
			(a < 0 && b < 0 && a < MaxInt/b) ||
			(a > 0 && b < 0 && b < MinInt/a) ||
			(a < 0 && b > 0 && a < MinInt/b) {
			return
		}
		printInt64(a * b)
	case "/":
		if b == 0 {
			printString("No division by 0")
			return
		}
		printInt64(a / b)
	case "%":
		if b == 0 {
			printString("No modulo by 0")
			return
		}
		printInt64(a % b)
	default:
		return
	}
}
