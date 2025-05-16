package main

import "fmt"

const NMAX = 100

type tabArray [NMAX]int

func inputData(t *tabArray, n *int, arrayName string) {
	var input int
	*n = 0
	fmt.Printf("Masukkan angka ke dalam array %s (akhiri dengan 0):\n", arrayName)
	for *n < NMAX {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		t[*n] = input
		*n++
	}
}

func cetakArray(t tabArray, n int) []int {
	array := make([]int, n)
	copy(array, t[:n])
	return array
}

func irisan(t1 tabArray, n1 int, t2 tabArray, n2 int) []int {

	exists := make(map[int]bool)
	for i := 0; i < n2; i++ {
		exists[t2[i]] = true
	}

	result := make([]int, 0, min(n1, n2))

	for i := 0; i < n1; i++ {
		if exists[t1[i]] {
			result = append(result, t1[i])
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var arrayA, arrayB tabArray
	var jmlhDataA, jmlhDataB int

	inputData(&arrayA, &jmlhDataA, "A")
	inputData(&arrayB, &jmlhDataB, "B")

	fmt.Println("Array A:", cetakArray(arrayA, jmlhDataA))
	fmt.Println("Array B:", cetakArray(arrayB, jmlhDataB))
	fmt.Println("Irisan array A dan B:", irisan(arrayA, jmlhDataA, arrayB, jmlhDataB))
}
