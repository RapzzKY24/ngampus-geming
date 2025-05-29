package main

import "fmt"

const NMAX = 1000

type tabBunga [NMAX]string

func bacaInputan(t *tabBunga, n *int) {
	var input string
	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan nama bunga ke-%d: ", i+1)
		fmt.Scan(&input)
		(*t)[i] = input
	}
}

func cetak(t tabBunga, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, t[i])
	}
}

func panjangBunga(flowers string) int {
	count := 0
	for _, flower := range flowers {
		if flower != '_' && flower != '.' {
			count++
		}
	}
	return count
}

func insertionSort(t *tabBunga, n int) {
	for i := 1; i < n; i++ {
		key := (*t)[i]
		keyLen := panjangBunga(key)
		j := i - 1
		for j >= 0 && panjangBunga((*t)[j]) > keyLen {
			(*t)[j+1] = (*t)[j]
			j--
		}
		(*t)[j+1] = key
	}
}

func main() {
	var data tabBunga
	var n int
	bacaInputan(&data, &n)
	insertionSort(&data, n)
	cetak(data, n)
}
