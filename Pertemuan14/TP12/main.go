package main

import "fmt"

const NMAX = 1000

type tabArray [NMAX]int

func inputData(t *tabArray, n *int) {
	var input int
	*n = 0
	fmt.Print("Masukkan data angka ke dalam array (diakhiri negatif)\n")
	for *n <= NMAX {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		t[*n] = input
		(*n)++
	}
}

func inserctionSort(t *tabArray, n int) {
	for i := 0; i < n; i++ {
		key := t[i]
		j := i - 1
		for j >= 0 && t[j] > key {
			t[j+1] = t[j]
			j--
		}
		t[j+1] = key
	}
}

func showArray(t tabArray, n int) []int {
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, t[i])
	}
	return array
}

func hitungJarak(t tabArray, n int) {
	jarak := t[1] - t[0]
	tetap := true

	for i := 1; i < n-1; i++ {
		currJarak := t[i+1] - t[i]
		if currJarak != jarak {
			tetap = false
			break
		}
	}

	if tetap {
		fmt.Printf("Jarak antar elemen array tetap, yaitu: %d\n", jarak)
	} else {
		fmt.Println("Jarak antar elemen array tidak tetap.")
	}
}

func main() {
	var data tabArray
	var jmlhData int
	inputData(&data, &jmlhData)
	fmt.Println("Data array sebelum di sorting: ", showArray(data, jmlhData))
	inserctionSort(&data, jmlhData)
	fmt.Println("Data array setelah di sorting: ", showArray(data, jmlhData))
	hitungJarak(data, jmlhData)
}
