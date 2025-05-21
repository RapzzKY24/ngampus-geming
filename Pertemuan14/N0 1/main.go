package main

import "fmt"

const NMAX = 1000

type tabArray [NMAX]int

func inputData(t *tabArray, n int, rumahKe string) {
	var input int
	fmt.Print("Masukkan total input:")
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
	}
}

func main() {

}
