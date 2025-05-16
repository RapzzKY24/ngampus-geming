package main

import "fmt"

const NMAX = 100

type ArrBalita [NMAX]float64

func cetakBalita(a *ArrBalita, n *int) {
	var inputBerat float64
	var jumlahAnak int
	*n = 0
	fmt.Print("Masukkan berat data banyak balita: ")
	fmt.Scan(&jumlahAnak)
	for *n < jumlahAnak && *n < NMAX {
		fmt.Printf("Masukkan berat balita ke-%d: ", *n+1)
		fmt.Scan(&inputBerat)
		a[*n] = inputBerat
		*n++
	}
}

func hitungMinMax(a ArrBalita, n int) (float64, float64) {
	max := a[0]
	min := a[0]
	for i := 0; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	return max, min
}

func rerata(a ArrBalita, n int) float64 {
	sum := 0.0
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	return float64(sum) / float64(n)
}

func main() {
	var jmlhData int
	var data ArrBalita
	cetakBalita(&data, &jmlhData)
	max, min := hitungMinMax(data, jmlhData)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rerata(data, jmlhData))
}
