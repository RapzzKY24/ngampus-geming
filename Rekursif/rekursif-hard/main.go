package main

import "fmt"

func cetakdigit(n int) {
	if n < 10 {
		fmt.Print(n)
	} else {
		cetakdigit(n / 10)
		fmt.Print("-", n%10)
	}
}

func jumlahDigit(n int) int {
	if n < 10 {
		return 1
	} else {
		return 1 + jumlahDigit(n/10)
	}
}

func faktor(i, N int) {
	if i <= N {
		// pengkondisian untuk mengecek nilai bilangan genap / ganjil
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
		faktor(i+1, N) // tuliskan fungsi rekursif pada baris ini sehingga output sesuai dengan yang diharapkan
	}
}

func lowercase(s string, n int) {
	/*{ I.S. Terdefinisi s sebagai input string kapital
	F.S. menampilkan string yang semua huruf kecil menggunakan fungsi rekursif*/
	if n > 0 {
		lowercase(s, n-1)              // Rekursif untuk karakter sebelumnya
		fmt.Print(string(s[n-1] + 32)) // Konversi karakter ke huruf kecil
	}
}

func cetakterbalikdigit(n int) {
	if n < 10 {
		fmt.Print(n)
	} else {
		fmt.Print(n%10, "-")
		cetakterbalikdigit(n / 10)
	}
}

func segitigaRecursiveGPT(baris, kolom, n int) {
	// Basis rekursi: jika sudah mencapai baris ke-n, berhenti
	if baris > n {
		return
	}

	// Cetak bintang sesuai kolom saat ini
	fmt.Print("*")

	// Pindah ke baris baru jika jumlah kolom sudah sesuai dengan nomor baris
	if kolom == baris {
		fmt.Println()
		segitigaRecursive(baris+1, 1, n) // Pindah ke baris berikutnya
	} else {
		segitigaRecursive(baris, kolom+1, n) // Lanjutkan mencetak di baris yang sama
	}
}

func segitigaRecursive(baris int, kolom int, n int) {
	if kolom <= n {
		if baris+kolom <= n { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan
			fmt.Print(" ")
		} else {
			fmt.Print("*")
		}
		segitigaRecursive(baris, kolom+1, n) // Rekursif untuk kolom berikutnya

	} else if baris < n {
		fmt.Println()
		segitigaRecursive(baris+1, 1, n) // Rekursif untuk baris berikutnya
	}
}

func printPolaBilangan(baris, kolom, n int) {
	if n >= baris {
		if n >= kolom {
			if baris == 1 || baris == n || kolom == 1 || kolom == n { // Pengecekan Pengulangan kolom dan baris dengan operasi perbandingan
				fmt.Print("*", " ")
			} else {
				fmt.Print("  ")
			}
			printPolaBilangan(baris, kolom+1, n) // Rekursif untuk kolom berikutnya
		} else {
			fmt.Println()
			printPolaBilangan(baris+1, 1, n) // Rekursif untuk baris berikutnya
		}
	}

}
func main() {
	faktor(1, 8)
}
