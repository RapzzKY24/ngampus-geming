package main

import "fmt"

type mahasiswa struct {
	nama, indeks string
	nilai        int
}

const NMAX int = 1024

type dataMahasiswa [NMAX]mahasiswa

func isiArray(himpunan *dataMahasiswa, n int) {
	var nama string
	var nilai int
	for i := 0; i < n; i++ {
		fmt.Scan(&nama, &nilai)
		var indeks string
		if nilai > 80 {
			indeks = "A"
		} else if nilai > 70 {
			indeks = "AB"
		} else if nilai > 65 {
			indeks = "B"
		} else if nilai > 60 {
			indeks = "BC"
		} else if nilai > 50 {
			indeks = "C"
		} else if nilai > 40 {
			indeks = "D"
		} else {
			indeks = "E"
		}
		(*himpunan)[i] = mahasiswa{
			nama:   nama,
			indeks: indeks,
			nilai:  nilai,
		}
	}
}

func insertionSort(himpunan *dataMahasiswa, n int) {
	for i := 1; i < n; i++ {
		key := (*himpunan)[i]
		j := i - 1
		for j >= 0 && (*himpunan)[j].nilai > key.nilai {
			(*himpunan)[j+1] = (*himpunan)[j]
			j--
		}
		(*himpunan)[j+1] = key
	}
}

func showArray(himpunan dataMahasiswa, n int, tampilIndeks string) {
	for i := 0; i < n; i++ {
		if himpunan[i].indeks == tampilIndeks {
			fmt.Printf("%s %d %s\n", himpunan[i].nama, himpunan[i].nilai, himpunan[i].indeks)
		}
	}
}

func main() {
	var n int
	var himpunan dataMahasiswa
	var tampilIndeks string

	fmt.Scan(&n)
	fmt.Scan(&tampilIndeks)

	isiArray(&himpunan, n)
	insertionSort(&himpunan, n)
	showArray(himpunan, n, tampilIndeks)
}
