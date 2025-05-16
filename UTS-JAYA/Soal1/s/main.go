package main

import "fmt"

// BISMILLAH

type Waktu struct {
	jam   int
	menit int
	detik int
}

func konversiDetik(w Waktu) int {
	return w.jam*3600 + w.menit*60 + w.detik //mengkoversi waktu ke detik agar bisa menghitung selisih pemakaian warnet
}

func konversiJam(totalDetik int) (int, int, int) {
	jam := totalDetik / 3600
	sisa := totalDetik % 3600
	menit := sisa / 60
	second := sisa % 60
	return jam, menit, second //mengembalikan ke bentuk semula
}

func inputData(w *Waktu, ket string) { //membuat inputan user
	fmt.Printf("Masukkan waktu %s\n", ket)
	fmt.Print("Masukkan jam: ")
	fmt.Scan(&w.jam)
	fmt.Print("Masukkan menit: ")
	fmt.Scan(&w.menit)
	fmt.Print("Masukkan detik: ")
	fmt.Scan(&w.detik)
}

func main() {
	var waktuMulai, waktuSelesai Waktu
	inputData(&waktuMulai, "Mulai")
	inputData(&waktuSelesai, "Selesai")
	totalWaktuMulai := konversiDetik(waktuMulai)
	totalWaktuSelesai := konversiDetik(waktuSelesai)
	selisih := totalWaktuSelesai - totalWaktuMulai

	if selisih < 0 {
		selisih += 24 * 3600
	}

	jam, menit, detik := konversiJam(selisih)
	fmt.Printf("%02d jam %02d menit %02d detik\n", jam, menit, detik)
}
