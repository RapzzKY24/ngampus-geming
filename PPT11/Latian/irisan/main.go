package main

import "fmt"

const CAPS = 100

type tabArray [CAPS]int

func inputData(t *tabArray, n *int, himpunan string) {
	*n = 0
	var input int
	fmt.Printf("Masukkan anggota himpunan %s (angka diakhiri 0)\n", himpunan)
	for *n <= CAPS {
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		t[*n] = input
		(*n)++
	}
}

func irisan(t1 tabArray, n1 int, t2 tabArray, n2 int) []int {
	var irisan []int
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			if t1[i] == t2[j] {
				irisan = append(irisan, t1[i])
			}
		}
	}
	return irisan
}

func main() {
	var himpunanA, himpunanB tabArray
	var jmlhDataA, jmlhDataB int
	inputData(&himpunanA, &jmlhDataA, "a")
	inputData(&himpunanB, &jmlhDataB, "b")

	fmt.Println("Hasil himpunan irisan dari a dan b adalah:", irisan(himpunanA, jmlhDataA, himpunanB, jmlhDataB))

}
