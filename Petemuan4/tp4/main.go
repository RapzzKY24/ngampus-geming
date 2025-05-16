package main

import "fmt"

func printAsterisRecursion(n int) {
	if n == 0 {
		return
	}
	printAsterisRecursion(n - 1)
	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)
	printAsterisRecursion(n)
}
