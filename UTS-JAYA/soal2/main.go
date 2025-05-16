package main

import "fmt"

//bismillah

const NMAX = 100

type tabInt [NMAX]int

func inputData(t *tabInt, n *int, nBil int) {
	var inputData int
	*n = 0
	fmt.Print("Masukkan jumlah bilangan yang ingin dimasukkan (nBil): ")
	fmt.Scan(&nBil)
	for *n < nBil && *n < NMAX {
		fmt.Printf("Masukkan data ke array %d: ", *n+1)
		fmt.Scan(&inputData)
		t[*n] = inputData
		*n++
	}
}

func cetak(t tabInt, n int) []int {
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, t[i])
	}
	return array
}

func jumlah(t tabInt, n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return sum
}

func average(t tabInt, n int) float64 {
	count := jumlah(t, n)
	return float64(count) / float64(n)
}

func diatasRataRata(t tabInt, n int) int {
	totalDiatasRata := 0
	for i := 0; i < n; i++ {
		if t[i] > int(average(t, n)) {
			totalDiatasRata++
		}
	}
	return totalDiatasRata
}

func main() {
	var jmlhData, nBil int
	var data tabInt
	inputData(&data, &jmlhData, nBil)
	cetakDatas := cetak(data, jmlhData)
	average := average(data, jmlhData)
	fmt.Printf("Rata-rata yang dihasilkan adalah: %.2f\n", average)
	fmt.Println("Data: ", cetakDatas)
	fmt.Println("Total diatas rata rata adalah: ", diatasRataRata(data, jmlhData))
}
