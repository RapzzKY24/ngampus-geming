package main

import "fmt"

const NMAX = 2022

type Nasabah struct {
	namaNasabah, namaBank, idNasabah string
	nomorRekening                    int
}

type tabNasabah [NMAX]Nasabah

func addNasabah(t *tabNasabah, n *int) {
	var idNasabah string
	var nomorRekening int
	var namaNasabah, namaBank string
	*n = 0
	for *n < NMAX {
		fmt.Print("Masukkan ID nasabah (ketik '0' untuk berhenti): ")
		fmt.Scan(&idNasabah)
		if idNasabah == "0" {
			break
		}
		fmt.Print("Masukkan nama nasabah: ")
		fmt.Scan(&namaNasabah)
		fmt.Print("Masukkan nomor rekening: ")
		fmt.Scan(&nomorRekening)
		fmt.Print("Masukkan nama bank: ")
		fmt.Scan(&namaBank)
		t[*n] = Nasabah{
			idNasabah:     idNasabah,
			nomorRekening: nomorRekening,
			namaNasabah:   namaNasabah,
			namaBank:      namaBank,
		}
		*n++
	}
}

func cetak(t tabNasabah, n int) {
	fmt.Println("Daftar Nasabah:")
	for i := 0; i < n; i++ {
		fmt.Printf("%-10s %-20s %-15d %-15s\n", t[i].idNasabah, t[i].namaNasabah, t[i].nomorRekening, t[i].namaBank)
	}
}

func searchBank(t tabNasabah, n int, namaBank string) {
	fmt.Printf("Nasabah dengan bank %s:\n", namaBank)
	for i := 0; i < n; i++ {
		if t[i].namaBank == namaBank {
			fmt.Printf("%-10s %-20s %-15d %-15s\n", t[i].idNasabah, t[i].namaNasabah, t[i].nomorRekening, t[i].namaBank)
		}
	}
}

func main() {
	var data tabNasabah
	var jmlhdata int
	var namaBank string

	addNasabah(&data, &jmlhdata)
	fmt.Println()

	cetak(data, jmlhdata)
	fmt.Println()

	fmt.Print("Masukkan nama bank yang ingin dicari: ")
	fmt.Scan(&namaBank)

	searchBank(data, jmlhdata, namaBank)
}
