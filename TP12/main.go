package main

import "fmt"

const NMAX = 100

type tabInt [NMAX]int

func inputData(t *tabInt, n *int) {
	var input int
	fmt.Print("Masukkan suara diakhiri dengan 0")
	for *n < NMAX {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		t[*n] = input
		(*n)++
	}
}

func cetakDataValid(t tabInt, n int) []int {
	dataTerbaca := []int{}
	for i := 0; i < n; i++ {
		dataTerbaca = append(dataTerbaca, t[i])
	}
	return dataTerbaca
}

func frekuensiDataValid(t tabInt, n int) int {
	frekuensi := 0
	for i := 0; i < n; i++ {
		frekuensi++
	}
	return frekuensi
}

func checkDataValid(t tabInt, n int) int {
	dataValid := 0
	for i := 0; i < n; i++ {
		if t[i] >= 1 && t[i] <= 15 {
			dataValid++
		}
	}
	return dataValid
}

func frekuensiWithSequentialSearch(t tabInt, n int, target int) int {
	fmt.Print("Masukkan target ingin dicari: ")
	fmt.Scan(&target)
	frekuensi := 0
	for i := 0; i < n; i++ {
		if t[i] == target {
			frekuensi++
		}
	}
	if frekuensi > 0 {
		fmt.Printf("Peserta %d mendapat %d suara\n", target, frekuensi)
	}
	return frekuensi
}

func main() {
	var jumlahData int
	var data tabInt
	var target int
	inputData(&data, &jumlahData)
	cetakDatas := cetakDataValid(data, jumlahData)
	checkDataValid := checkDataValid(data, jumlahData)
	checkFrekuensiValid := frekuensiDataValid(data, jumlahData)
	fmt.Println("Jumlah suara yang terbaca :", cetakDatas)
	fmt.Println("jumlah suuara yang terbaca :", checkFrekuensiValid)
	fmt.Println("Jumlah suara valid :", checkDataValid)
	frekuensiWithSequentialSearch(data, jumlahData, target)
}
