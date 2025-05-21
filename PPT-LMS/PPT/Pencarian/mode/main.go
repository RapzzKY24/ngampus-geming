package main

import "fmt"

const NMAX = 1000

type tabArray [NMAX]int

func frekuensi(t tabArray, n int, x int) int {
	count := 0
	for i := 0; i < n; i++ {
		if x == t[i] {
			count++
		}
	}
	return count
}

func membuatTabel(t1 tabArray, t2 *tabArray, n int, m int) int {
	count := 0
	m = 10
	for i := 0; i <= m; i++ {
		t2[i] = frekuensi(t1, n, i)
		i++
	}
	return count
}

func modus(t tabArray, n int) int {
	freq := tabArray{}
	maxFreq := 0
	modus := -1

	for i := 0; i < n; i++ {
		freq[t[i]]++
		if freq[t[i]] > maxFreq || (freq[t[i]] == maxFreq && t[i] < modus) {
			maxFreq = freq[t[i]]
			modus = t[i]
		}
	}

	return modus
}

func main() {
	var t tabArray
	var n int

	println("Masukkan jumlah elemen:")
	fmt.Scan(&n)

	println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}

	println("Array yang dimasukkan adalah:")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], " ")
	}
	println()
	println("Frekuensi elemen dalam array:")
	for i := 0; i < n; i++ {
		fmt.Print(t[i], ":", frekuensi(t, n, t[i]), " ")
	}
	println()

	println("Tabel frekuensi:")
	var t2 tabArray
	membuatTabel(t, &t2, n, 10)
	for i := 0; i < 10; i++ {
		fmt.Print(i, ":", t2[i], " ")
	}
	println()

	result := modus(t, n)

	println("Modus dari array adalah:", result)
}
