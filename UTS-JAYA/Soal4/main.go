package main

import "fmt"

type TiketPesawat struct {
	maskapai string
	kelas    string
	harga    int
}

const NMAX = 100

type tabTiket [NMAX]TiketPesawat

func inputNamaMaskapai(t *tabTiket, n *int) {
	*n = 0
	var maskapai, kelas string
	var harga int
	for *n < NMAX {
		fmt.Print("Masukkan maskapai: ")
		fmt.Scan(&maskapai)
		fmt.Print("Masukkan kelas: ")
		fmt.Scan(&kelas)
		fmt.Print("Masukkan harga: ")
		fmt.Scan(&harga)
		if maskapai == "none" && kelas == "none" && harga == 0 {
			break
		}
		t[*n] = TiketPesawat{
			maskapai: maskapai,
			kelas:    kelas,
			harga:    harga,
		}
		*n++
	}
}

func cetak(t tabTiket, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Maskapai: %s, Kelas: %s, Harga: %d\n", t[i].maskapai, t[i].kelas, t[i].harga)
	}
}

func cariTiketTermahal(t tabTiket, n int) TiketPesawat {
	if n == 0 {
		return TiketPesawat{}
	}
	termahal := t[0]
	for i := 1; i < n; i++ {
		if t[i].harga > termahal.harga {
			termahal = t[i]
		}
	}
	return termahal
}

func main() {
	var tiket tabTiket
	var n int

	inputNamaMaskapai(&tiket, &n)
	fmt.Println()
	cetak(tiket, n)

	if n > 0 {
		termahal := cariTiketTermahal(tiket, n)
		fmt.Println("Tiket termahal:")
		fmt.Printf("Maskapai: %s, Kelas: %s, Harga: %d\n", termahal.maskapai, termahal.kelas, termahal.harga)
	} else {
		fmt.Println("Tidak ada tiket yang dimasukkan.")
	}
}
