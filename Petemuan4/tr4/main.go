package main

import "fmt"

func recursionFindFactor(n, i int) {
	if n < i {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	recursionFindFactor(n, i+1)
}

func adaBilangan(n, m int) bool {
	for n > 0 {
		digit := n % 10
		if digit == m {
			return true
		}
		n /= 10
	}
	return false
}

func fib(n int) int {
	if n <= 3 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	recursionFindFactor(8, 1)
}
