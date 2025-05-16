package main

import "fmt"

const caps = 1000

type tabInt [caps]int

func inputData(t *tabInt, n *int, himpunan string) {
	var inputDatas int
	*n = 0
	fmt.Printf("Masukkan angka ke dalam array %s (akhiri dengan 0):\n", himpunan)
	for *n <= caps {
		fmt.Scan(&inputDatas)

		if inputDatas == 0 {
			break
		}
		t[*n] = inputDatas
		(*n)++
	}
}

func copy(t tabInt, n int) []int {
	var gabungan []int
	for i := 0; i < n; i++ {
		gabungan = append(gabungan, t[i])
	}
	return gabungan
}

func sudahAda(slice []int, val int) bool {
	for _, num := range slice {
		if num == val {
			return true
		}
	}
	return false
}

func gabungan(t1, t2 tabInt, n1, n2 int) []int {
	copyA := copy(t1, n1)
	for j := 0; j < n2; j++ {
		if !sudahAda(copyA, t2[j]) {
			copyA = append(copyA, t2[j])
		}
	}
	return copyA
}

func main() {
	var data1, data2 tabInt
	var jmlhData1, jmlhData2 int
	inputData(&data1, &jmlhData1, "a")
	inputData(&data2, &jmlhData2, "b")
	fmt.Println(copy(data1, jmlhData1))
	fmt.Println()
	fmt.Println(gabungan(data1, data2, jmlhData1, jmlhData2))
}
