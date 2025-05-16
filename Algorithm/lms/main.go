package main

import "fmt"

const NMAX = 100

type tabArray [NMAX]int

func inputDat(t *tabArray, n *int, arrayName string) {
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
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, t[i])
	}
	return array
}

func irisan(t1 tabArray, n1 int, t2 tabArray, n2 int) []int {
	var irisan []int
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if t1[i] == t2[j] {
				irisan = append(irisan, t1[i])
				break
			}
		}
	}
	return irisan
}

func main() {
	var arrayA, arrayB tabArray
	var jmlhDataA, jmlhDataB int

	inputDat(&arrayA, &jmlhDataA, "A")

	inputDat(&arrayB, &jmlhDataB, "B")

	fmt.Println("Array A: ", cetakArray(arrayA, jmlhDataA))

	fmt.Println("Array B: ", cetakArray(arrayB, jmlhDataB))

	fmt.Println("Irisan array A dan B:", irisan(arrayA, jmlhDataA, arrayB, jmlhDataB))
}
