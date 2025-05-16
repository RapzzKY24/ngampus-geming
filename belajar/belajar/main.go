package main

import "fmt"

const NMAX = 1000

type Lulusan struct {
	nama, nim string
	semester  int
	eprt, ipk float64
}

type tabLulus [NMAX - 1]Lulusan

func isiData(t *tabLulus, n *int) {
	var inputNIM string
	fmt.Println("Masukkan NIM (ketik 'none' untuk berhenti): ")
	fmt.Scan(&inputNIM)
	for inputNIM != "none" && *n < NMAX {
		t[*n].nim = inputNIM
		fmt.Println("Masukkan Nama: ")
		fmt.Scan(&t[*n].nama)
		fmt.Println("Masukkan nilai EPRT: ")
		fmt.Scan(&t[*n].eprt)
		fmt.Println("Masukkan semester: ")
		fmt.Scan(&t[*n].semester)
		fmt.Println("Masukkan IPK: ")
		fmt.Scan(&t[*n].ipk)
		(*n)++
		fmt.Println("Masukkan NIM (ketik 'none' untuk berhenti)")
		fmt.Scan(&inputNIM)
	}
}

func maxEPRT(t tabLulus, n int) float64 {
	max := t[0].eprt
	for i := 1; i < n; i++ {
		if t[i].eprt > max {
			max = t[i].eprt
		}
	}
	return max
}

func maxIPK(t tabLulus, n int) float64 {
	max := t[0].ipk
	for i := 1; i < n; i++ {
		if t[i].ipk > max {
			max = t[i].ipk
		}
	}
	return max
}

func rataRata(t tabLulus, n int) float64 {
	if n == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i].semester
	}
	return float64(sum) / float64(n)
}

func main() {
	var data tabLulus
	var jmlData int

	fmt.Println("Mulai memasukkan data lulusan:")
	isiData(&data, &jmlData)

	if jmlData > 0 {
		fmt.Printf("Nilai EPRT tertinggi: %.2f\n", maxEPRT(data, jmlData))
		fmt.Printf("Nilai IPK tertinggi: %.2f\n", maxIPK(data, jmlData))
		fmt.Printf("Rata-rata semester: %.2f\n", rataRata(data, jmlData))
	} else {
		fmt.Println("Tidak ada data yang dimasukkan.")
	}
}
