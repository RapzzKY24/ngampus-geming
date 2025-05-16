package main

import "fmt"

const NMAX = 1000000

var data [NMAX]int

func isiArray(n int) {
	var input int
	fmt.Println("Masukkan bilangan array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		data[i] = input // mengisi variabel data dengan inputan user
	}
}

func posisi(n, k int) int {
	for i := 0; i < n; i++ {
		if data[i] == k { // sequential search
			return i // return index
		}
	}
	return -1 // ga nemu
}

func main() {
	var n, k int
	fmt.Print("Masukkan jumlah isi array: ")
	fmt.Scan(&n)
	fmt.Print("Masukkan angka yang ingin dicari: ")
	fmt.Scan(&k)
	isiArray(n)
	result := posisi(n, k)
	if result == -1 {
		fmt.Printf("Posisi %d tidak ditemukan dalam array maka di return %d", k, result) // jika hasil -1 maka return gaada
	} else {
		fmt.Printf("Posisi %d ditemukan di array pada indeks ke- %d", k, result) // jika ada return result
	}
}
