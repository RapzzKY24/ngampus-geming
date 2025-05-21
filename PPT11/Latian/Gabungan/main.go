package main

import "fmt"

const CAPS = 1000

type tabArray [CAPS]int

func inputData(t *tabArray, n *int, namaHimpunan string) {
	var input int
	*n = 0
	fmt.Printf("Masukkan data himpunan %s ke array:\n", namaHimpunan)
	for *n <= CAPS {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		t[*n] = input
		(*n)++
	}
}

func copy(t tabArray, n int) []int {
	var gabungan []int
	for i := 0; i < n; i++ {
		gabungan = append(gabungan, t[i])
	}
	return gabungan
}

func sudahAda(numbers []int, val int) bool {
	for _, num := range numbers {
		if num == val {
			return true
		}
	}
	return false
}

func checkGabungan(t1 tabArray, n1 int, t2 tabArray, n2 int) []int {
	union := copy(t1, n1)
	for i := 0; i < n2; i++ {
		if !sudahAda(union, t2[i]) {
			union = append(union, t2[i])
		}
	}
	return union
}

func main() {
	var himpunanA, himpunanB tabArray
	var jmlhDataA, jmlhDataB int
	inputData(&himpunanA, &jmlhDataA, "a")
	inputData(&himpunanB, &jmlhDataB, "b")
	fmt.Println(checkGabungan(himpunanA, jmlhDataA, himpunanB, jmlhDataB))
}
