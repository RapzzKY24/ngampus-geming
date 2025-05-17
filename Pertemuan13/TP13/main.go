package main

import "fmt"

const NMAX = 100

type tabArray [NMAX]int

func isiData(t *tabArray, n *int) {
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(n)
	var input int
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan data ke %d: ", i+1)
		fmt.Scan(&input)
		t[i] = input
	}
	fmt.Println(t[:*n])
}

func binerySearch(t tabArray, n int, target int) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if t[mid] == target {
			return mid
		} else if t[mid] < target {
			right = mid + 1
		} else {
			left = mid - 1
		}
	}
	return -1
}

func main() {
	var data tabArray
	var jumlahData, target int
	isiData(&data, &jumlahData)
	fmt.Print("Masukkan angka ingin dicari: ")
	fmt.Scan(&target)
	fmt.Println(binerySearch(data, jumlahData, target))
}
