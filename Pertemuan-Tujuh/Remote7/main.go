package main

import (
	"fmt"
	"math"
)

const NMAX = 1000

type tabArray [NMAX - 1]int

func bacaData(t *tabArray, n *int) {
	var input int
	*n = 0
	for *n < NMAX {
		fmt.Printf("Masukkan angka ke-%d: ", *n+1)
		fmt.Scan(&input)

		if input == 0 {
			break
		}

		t[*n] = input
		*n++
	}
}

func cetak(t tabArray, n int) []int {
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, t[i])
	}
	return array
}

func idxGanjil(t tabArray, n int) {
	for i := 0; i < n; i++ {
		if i%2 != 0 {
			fmt.Println(t[i])
		}
	}
}
func idxGenap(t tabArray, n int) {
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Println(t[i])
		}
	}
}

func kelipatanX(t tabArray, n int, x int) []int {
	var arrayKelipatan []int
	for i := 0; i < n; i++ {
		if i%x == 0 {
			arrayKelipatan = append(arrayKelipatan, t[i])
		}
	}
	return arrayKelipatan
}

func removeData(t *tabArray, n *int, targetRemove int) []int {
	j := 0
	for i := 0; i < *n; i++ {
		if t[i] != targetRemove {
			t[j] = t[i]
			j++
		}
	}
	*n = j
	return cetak(*t, *n)
}

func rataRata(t tabArray, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return float64(sum) / float64(n)
}

func standarDeviasi(t tabArray, n int) float64 {
	if n == 0 {
		return 0
	}

	mean := rataRata(t, n)
	var sum float64

	for i := 0; i < n; i++ {
		diff := float64(t[i]) - mean
		sum += diff * diff
	}

	variance := sum / float64(n)
	return math.Sqrt(variance)
}

func frekuensiFromHashMap(t tabArray, n int) {
	frequencyMap := make(map[int]int)

	for i := 0; i < n; i++ {
		frequencyMap[t[i]]++
	}

	for key, value := range frequencyMap {
		fmt.Printf("Angka %d muncul sebanyak %d kali\n", key, value)
	}
}

func main() {
	var data tabArray
	var jumlahData int
	var choice int

	for {
		fmt.Println("========================================")
		fmt.Println("         PROGRAM PENGOLAH DATA          ")
		fmt.Println("========================================")
		fmt.Println(" 1.  Input Data")
		fmt.Println(" 2.  Cetak Data")
		fmt.Println(" 3.  Tampilkan Data di Indeks Ganjil")
		fmt.Println(" 4.  Tampilkan Data di Indeks Genap")
		fmt.Println(" 5.  Tampilkan Kelipatan X")
		fmt.Println(" 6.  Hapus Data")
		fmt.Println(" 7.  Hitung Rata-rata")
		fmt.Println(" 8.  Hitung Standar Deviasi")
		fmt.Println(" 9.  Frekuensi Nilai")
		fmt.Println(" 10. Keluar")
		fmt.Println("========================================")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			bacaData(&data, &jumlahData)
		case 2:
			fmt.Println("Data yang dimasukkan:")
			fmt.Println(cetak(data, jumlahData))
		case 3:
			fmt.Println("Data pada indeks ganjil:")
			idxGanjil(data, jumlahData)
		case 4:
			fmt.Println("Data pada indeks genap:")
			idxGenap(data, jumlahData)
		case 5:
			var x int
			fmt.Print("Masukkan nilai X untuk mencari kelipatan: ")
			fmt.Scan(&x)
			fmt.Printf("Data pada indeks kelipatan %d:\n", x)
			fmt.Println(kelipatanX(data, jumlahData, x))
		case 6:
			var targetRemove int
			fmt.Print("Masukkan angka yang ingin dihapus: ")
			fmt.Scan(&targetRemove)
			fmt.Println("Data setelah penghapusan:")
			fmt.Println(removeData(&data, &jumlahData, targetRemove))
		case 7:
			fmt.Printf("Rata-rata data: %.2f\n", rataRata(data, jumlahData))
		case 8:
			fmt.Printf("Standar deviasi data: %.2f\n", standarDeviasi(data, jumlahData))
		case 9:
			fmt.Println("Frekuensi kemunculan nilai:")
			frekuensiFromHashMap(data, jumlahData)
		case 10:
			fmt.Println("Keluar dari program. Terima kasih.")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}

}
