package main

import "fmt"

const NMAX = 10

type tabInt [NMAX]int

func bacaData(A *tabInt, n *int) {
	var input int
	*n = 0
	for *n < NMAX {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&input)

		if input == 0 {
			break
		}
		if input < 0 {
			input *= -1
		}
		A[*n] = input
		*n++
	}
}

func cetak(A tabInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i], "\n")
	}
}

func jumlah(A tabInt, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += A[i]
	}
	return sum
}

func rata(A tabInt, n int) float64 {
	sum := jumlah(A, n)
	return float64(sum) / float64(n)
}

func main() {
	var data tabInt
	var jmlData int

	bacaData(&data, &jmlData)
	cetak(data, jmlData)
	fmt.Println("Jumlah keseluruhan : ", jumlah(data, jmlData))
	fmt.Println("Hasil rata :", rata(data, jmlData))
}
