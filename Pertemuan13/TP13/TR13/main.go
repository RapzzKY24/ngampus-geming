package main

import "fmt"

const NMAX = 1000000

var array [NMAX]int

func isiArray(n int) {
	var input int
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan angka ke-%d :", i+1)
		fmt.Scan(&input)
		array[i] = input
	}
}

func posisi(n, k int) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == k {
			return mid
		} else if array[mid] < k {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var n, k int
	fmt.Print("Masukkan jumlah data array: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan target yang ingin dicari: ")
	fmt.Scan(&k)
	isiArray(n)
	result := posisi(n, k)
	if result == -1 {
		fmt.Printf("Posisi %d tidak ditemukan dalam array maka di return %d", k, result) // jika hasil -1 maka return gaada
	} else {
		fmt.Printf("Posisi %d ditemukan di array pada indeks ke-%d", k, result) // jika ada return result
	}
}
