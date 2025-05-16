package main

import "fmt"

const GOL = 1000

type tabGol [GOL]int

func inputData(t *tabGol, n *int, tim string) {
	*n = 0
	var inputGol int
	for *n < GOL {
		fmt.Printf("Masukkan gol untuk tim %s pada pertandingan %d: ", tim, *n+1)
		fmt.Scan(&inputGol)
		if inputGol == -1 {
			break
		}
		t[*n] = inputGol
		(*n)++
	}
}

func cetak(t tabGol, n int) []int {
	var arrayGol []int
	for i := 0; i < n; i++ {
		arrayGol = append(arrayGol, t[i])
	}
	return arrayGol
}

func menentukanKemenagan(t1 tabGol, n1 int, t2 tabGol, n2 int) (int, int) {
	if n1 != n2 {
		fmt.Println("Jumlah pertandingan tidak sama, tidak bisa menentukan pemenang.")
		return 0, 0
	}

	menangA := 0
	menangB := 0
	for i := 0; i < n1; i++ {
		if t1[i] > t2[i] {
			menangA++
		} else if t1[i] < t2[i] {
			menangB++
		} else {
			fmt.Printf("Pertandingan %d berakhir seri dengan skor %d vs %d\n", i+1, t1[i], t2[i])
		}
	}
	return menangA, menangB
}

func rerataKemenangan(t1, t2 tabGol, n1, n2 int) (float64, float64) {
	if n1 == 0 || n2 == 0 {
		fmt.Println("Jumlah pertandingan tidak valid untuk menghitung rata-rata.")
		return 0, 0
	}

	menangA, menangB := menentukanKemenagan(t1, n1, t2, n2)
	rataKemenanganA := float64(menangA) / float64(n1)
	rataKemenanganB := float64(menangB) / float64(n2)
	return rataKemenanganA, rataKemenanganB
}

func main() {
	var data1, data2 tabGol
	var jmlhdata1, jmlhdata2 int

	fmt.Println("Input data untuk tim A:")
	inputData(&data1, &jmlhdata1, "A")
	fmt.Println("\nInput data untuk tim B:")
	inputData(&data2, &jmlhdata2, "B")

	fmt.Println("\nHasil gol tim A:", cetak(data1, jmlhdata1))
	fmt.Println("Hasil gol tim B:", cetak(data2, jmlhdata2))

	menangA, menangB := menentukanKemenagan(data1, jmlhdata1, data2, jmlhdata2)
	fmt.Printf("\nTim A menang %d kali, Tim B menang %d kali.\n", menangA, menangB)

	rataTimA, rataTimB := rerataKemenangan(data1, data2, jmlhdata1, jmlhdata2)
	fmt.Printf("\nTim A menang %.2f kali, Tim B menang %.2f kali.\n", rataTimA, rataTimB)
}
