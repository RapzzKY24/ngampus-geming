package main

import "fmt"

const CAPS = 100

type tabArray [CAPS]int

func inputData(t *tabArray, n *int) {
	var input int
	*n = 0
	fmt.Print("Masukkan anggota himpunan\n")
	for *n <= CAPS {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		t[*n] = input
		(*n)++
	}
}

func frekuensi(t tabArray, n int, x int) int {
	count := 0
	i := 0
	for i < n {
		if x == t[i] {
			count += 1
		}
		i++
	}
	return count
}

func tableFrekuensi(t1 tabArray, n int, t2 *tabArray) {
	m := 10 // asumsikan maksimal frekuensi 1 - 10
	i := 0
	for i <= m {
		t2[i] = frekuensi(t1, n, i)
		i++
	}
}

func checkValid(t tabArray, n int) string {
	var tabFreg tabArray
	tableFrekuensi(t, n, &tabFreg)
	for i := 0; i < n; i++ {
		if frekuensi(t, n, t[i]) != 1 {
			return "Tidak Valid"
		}
	}
	return "Valid"
}

func main() {
	var himpunan tabArray
	var jmlhData int

	inputData(&himpunan, &jmlhData)
	fmt.Println(checkValid(himpunan, jmlhData))

}
