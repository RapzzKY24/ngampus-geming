package main

import "fmt"

type Pasukan struct {
	namaJendral, jenis string
	jumlah             int
}

const NMAX = 1000

type tabPasukan [NMAX]Pasukan

func InputData(p *tabPasukan, n *int) {
	var inputNamaJendral, inputJenis string
	var inputJumlah int
	*n = 0
	for *n < NMAX {
		fmt.Print("Masukkan nama jendral: ")
		fmt.Scan(&inputNamaJendral)
		fmt.Print("Masukkan jenis jendral: ")
		fmt.Scan(&inputJenis)
		fmt.Print("Masukkan jumlah jendral: ")
		fmt.Scan(&inputJumlah)

		if inputNamaJendral == "none" && inputJenis == "none" && inputJumlah == 0 {
			break
		}
		p[*n] = Pasukan{
			namaJendral: inputNamaJendral,
			jenis:       inputJenis,
			jumlah:      inputJumlah,
		}
		*(n)++
	}
}

func cetak(p tabPasukan, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%-10s %-20s %-15d\n", p[i].namaJendral, p[i].jenis, p[i].jumlah)
	}
}

func salinArray(p tabPasukan, n int) tabPasukan {
	var arrayBaru tabPasukan
	for i := 0; i < n; i++ {
		arrayBaru[i] = p[i]
	}
	return arrayBaru
}

func cariTiketTermahal(p tabPasukan, n int) Pasukan {
	if n == 0 {
		return Pasukan{}
	}

	termahal := p[0]
	for i := 1; i < n; i++ {
		if p[i].jumlah > termahal.jumlah {
			termahal = p[i]
		}
	}
	return termahal
}

func main() {
	var data tabPasukan
	var jmlhData int
	InputData(&data, &jmlhData)

	arrayBaru := salinArray(data, jmlhData)

	fmt.Println("Data asli:")
	cetak(data, jmlhData)

	fmt.Println("\nData di array baru:")
	cetak(arrayBaru, jmlhData)

	termahal := cariTiketTermahal(data, jmlhData)
	fmt.Println("\nTiket termahal:")
	fmt.Printf("%-10s %-20s %-15d\n", termahal.namaJendral, termahal.jenis, termahal.jumlah)
}
