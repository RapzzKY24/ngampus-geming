package main

import "fmt"

type barang struct {
	nama  string
	harga int
}

const NMAX int = 1024

type dataBarang [NMAX]barang

func isiArray(himpunan *dataBarang, n int) {
	var nama string
	var harga int
	for i := 0; i < n; i++ {
		fmt.Scan(&nama, &harga)
		(*himpunan)[i] = barang{
			nama:  nama,
			harga: harga,
		}
	}
}

func insertionSort(himpunan *dataBarang, n int) {
	for i := 0; i < n; i++ {
		key := (*himpunan)[i]
		j := i - 1
		for j >= 0 && (*himpunan)[j].harga < key.harga {
			(*himpunan)[j+1] = (*himpunan)[j]
			j--
		}
		(*himpunan)[j+1] = key
	}
}

func showArray(himpunan dataBarang, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].harga)
	}
}

func main() {
	var himpunan dataBarang
	var n int
	fmt.Scanln(&n)
	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n)
}
