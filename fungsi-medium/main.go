package main

import "fmt"

// Buat fungsi saja
// Buat fungsi saja

// func segitigaSamaSisi(s1, s2, s3 float64) bool {
// 	return s1 > 0 && s2 > 0 && s3 > 0 && s1 == s2 && s2 == s3
// 	// Mengembalikan nilai boolean jika 3 bilangan real dapat
// 	// membentuk segitiga sama sisi
// }

// implementasikan fungsinya saja

// Buatlah fungsi saja

// func uangSaku(kota string, bil int) int {
// 	if kota == "A" {
// 		return bil * 500000
// 	} else if kota == "B" {
// 		return bil * 350000
// 	}
// 	return bil * 250000
// }

// implementasikan fungsinya saja

// // implementasikan fungsinya saja

// func pertemuan(x, y int) int {
// 	/* mengembalikan jumlah pertemuan sesuai ketentuan soal */
// 	count := 0
// 	for i := 0; i < 365; i++ {
// 		if i%x == 0 && i%y != 0 {
// 			count++
// 		}
// 	}
// 	return count
// }

// func jumlahLetusan(x, y, n int) int {
// 	count := 0
// 	tahun := x

// 	for tahun <= n {
// 		count++
// 		if count%2 == 1 {
// 			tahun += x
// 		} else {
// 			tahun += y
// 		}
// 	}

// 	return count
// }

// func jumlahGanjil(bil1, bil2 int) int {
// 	/* mengembalikan penjumlahan bilangan ganjil dalam interval masukan */
// 	count := 0
// 	for i := bil1; i <= bil2; i++ {
// 		if i%2 != 0 {
// 			count += i
// 		}
// 	}
// 	return count
// }

//buatlah fungsinya saja

func cariMedian(a, b, c int) int {
	/*  function mengembalikan nilai median */
	max := a
	min := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return (max + min) / 2
}

func diskon(totalBelanja int) int {
	// Mengembalikan 1 jika totalBelanja habis dibagi 100 atau 4
	if totalBelanja%100 == 0 || totalBelanja%4 == 0 {
		return 1
	}
	return 0
}

func cashback(totalBelanja int) int {
	// Mengembalikan 1 jika totalBelanja habis dibagi 100 atau 3
	if totalBelanja%100 == 0 || totalBelanja%3 == 0 {
		return 1
	}
	return 0
}
func main() {
	// Data dari gambar
	inputs := [][]int{
		{68900, 66300, 42420, 83760, 79610, 66340, 49690}, // 7 pembeli
		{57960, 71130},                      // 2 pembeli
		{93100, 54500, 72200, 82900, 61600}, // 5 pembeli
	}

	// Uji tiap kasus
	for _, belanjaan := range inputs {
		diskonTotal := 0
		cashbackTotal := 0
		for _, total := range belanjaan {
			diskonTotal += diskon(total)
			cashbackTotal += cashback(total)
		}
		fmt.Println(diskonTotal, cashbackTotal)
	}
}
