package main

import "fmt"

func bilGanjil(n int, start int) {
	if start > n {
		return
	}
	if start%2 != 0 {
		fmt.Println(start)
	}
	bilGanjil(n, start+1)
}
func prima(n int) string {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	if count == 2 {
		return "Prima"

	} else {
		return "Not prima"
	}
}

func faktorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func fibonnaci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}

func xPangkatY(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * xPangkatY(x, y-1)
}

func asteris(n int) {
	if n == 0 {
		return
	}
	asteris(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	// var bil int
	// fmt.Print("Masukkan bilangan bulat: ")
	// fmt.Scan(&bil)
	// bilGanjil(bil, 1)

	// var n int
	// fmt.Print("Masukkan bilangan: ")
	// fmt.Scan(&n)
	// FindFactor(n, 1)

	// var fibo int
	// fmt.Print("Masukkn sebuah bilangan: ")
	// fmt.Scan(&fibo)
	// fmt.Println(fibonnaci(fibo))

	// var x int
	// fmt.Print("Masukkan jumlah asteris yang ingin dicetak: ")
	// fmt.Scan(&x)
	// asteris(x)
	fmt.Println(prima(5))

	// var bilangan1, bilangan2 int
	// fmt.Print("Masukkan bilangan 1: ")
	// fmt.Scan(&bilangan1)
	// fmt.Print("Masukkan bilangan 2: ")
	// fmt.Scan(&bilangan2)
	// fmt.Println(xPangkatY(bilangan1, bilangan2))
}
