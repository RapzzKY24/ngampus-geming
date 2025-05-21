package main

import "fmt"

type pemain struct {
	nama        string
	point       int
	panjangNama int
}

const NMAX int = 1024

type dataPemain [NMAX]pemain

func isiArray(himpunan *dataPemain, n int) {
	var nama string
	var point int
	for i := 0; i < n; i++ {
		fmt.Scanln(&nama, &point)
		(*himpunan)[i] = pemain{
			nama:        nama,
			point:       point,
			panjangNama: len(nama),
		}
	}
}

func selectionSort(himpunan *dataPemain, n int) {
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if (*himpunan)[j].panjangNama > (*himpunan)[maxIdx].panjangNama {
				maxIdx = j
			}
		}
		if maxIdx != i {
			(*himpunan)[i], (*himpunan)[maxIdx] = (*himpunan)[maxIdx], (*himpunan)[i]
		}
	}
}

func showArray(himpunan dataPemain, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(himpunan[i].nama, himpunan[i].point)
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	var himpunan dataPemain
	isiArray(&himpunan, n)
	selectionSort(&himpunan, n)
	showArray(himpunan, n)
}
