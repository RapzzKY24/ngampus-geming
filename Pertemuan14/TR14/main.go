package main

import "fmt"

type Rumah struct {
	Jumlah int
	Nomor  []int
}

func main() {
	var jumlahPerumahan int
	fmt.Scanln(&jumlahPerumahan)

	perumahan := make([]Rumah, jumlahPerumahan)
	for i := 0; i < jumlahPerumahan; i++ {
		var jumlahRumah int
		fmt.Scan(&jumlahRumah)
		perumahan[i].Jumlah = jumlahRumah
		perumahan[i].Nomor = inputDanUrutkanNomor(jumlahRumah)
	}

	fmt.Println("\n\nHasil:")
	for _, r := range perumahan {
		for _, nomor := range r.Nomor {
			fmt.Print(nomor, " ")
		}
		fmt.Println()
	}
}
func inputDanUrutkanNomor(jumlah int) []int {
	ganjil := []int{}
	genap := []int{}
	for i := 0; i < jumlah; i++ {
		var nomor int
		fmt.Scan(&nomor)
		if nomor%2 == 0 {
			genap = append(genap, nomor)
		} else {
			ganjil = append(ganjil, nomor)
		}
	}
	urutAscending(ganjil)
	urutDescending(genap)
	return append(ganjil, genap...)
}

func urutAscending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func urutDescending(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}
