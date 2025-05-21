package main

import "fmt"

const NMAX int = 37

type tHimpunan struct {
	anggota [NMAX]int
	panjang int
}

func bacaMasukan(h *tHimpunan) {
	var input int
	count := 0
	for i := 0; i < h.panjang; i++ {
		fmt.Scan(&input)
		if ada(*h, input) {
			break
		}
		h.anggota[count] = input
		count++
	}
	h.panjang = count
}

func ada(set tHimpunan, x int) bool {
	for i := 0; i < set.panjang; i++ {
		if set.anggota[i] == x {
			return true
		}
	}
	return false
}

func insertionSort(set *tHimpunan) {
	for i := 1; i < set.panjang; i++ {
		key := set.anggota[i]
		j := i - 1
		for j >= 0 && set.anggota[j] > key {
			set.anggota[j+1] = set.anggota[j]
			j--
		}
		set.anggota[j+1] = key
	}
}

func sama(set1, set2 tHimpunan) bool {
	if set1.panjang != set2.panjang {
		return false
	}
	for i := 0; i < set1.panjang; i++ {
		if !ada(set2, set1.anggota[i]) {
			return false
		}
	}
	return true
}

func main() {
	var himpunan1, himpunan2 tHimpunan
	fmt.Scanln(&himpunan1.panjang)
	bacaMasukan(&himpunan1)
	insertionSort(&himpunan1)

	fmt.Scanln(&himpunan2.panjang)
	bacaMasukan(&himpunan2)
	insertionSort(&himpunan2)

	if sama(himpunan1, himpunan2) {
		fmt.Println("Sama")
	} else {
		fmt.Println("Tidak Sama")
	}
}
