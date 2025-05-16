package main

import "fmt"

func kali(n int) int {
	/*{Terdefinisi sebuah bilangan bulat positif n.
	  Fungsi mengembalikan hasil penjumlahan 1, 2, 3, ...., n. }*/
	if n == 1 {
		return n
	} else {
		return n * kali(n-1)
	}
}

func lowerUpper(s string, n int) {
	/*{ I.S. Terdefinisi s sebagai input string dengan variasi case
	F.S. menampilkan string yang huruf kecilnya sudah diubah menjadi huruf kapital dan sebaliknya*/

	if n > 1 { // gunakan operator yang tepat
		return // masukan fungsi rekursif pada barus ini
	}
	if s[n-1] < 97 {
		fmt.Print(string(s[n] + 32)) // tampilkan string

	} else {
		fmt.Print(string(s[n] - 32)) // tampilkan string
	}
}

func recPrint(n int, s string) {
	/*I.S. Terdefinisi nilai bilangan bulat sebanyak n.
	  F.S.  menampilkan string sebanyak n kali*/

	if n == 1 {
		fmt.Println(s)
	} else {
		fmt.Println(s)
		recPrint(n-1, s)
	}
}

func kakiHewan(n int) int {
	if n == 0 {
		return 0
	}

	if n%2 == 1 {
		return 2 + kakiHewan(n-1)
	}
	return 3 + kakiHewan(n-1)
}

func lowercase(s string, n int) {
	/* { I.S. Terdefinisi s sebagai input string kapital
	   F.S. Menampilkan string yang semua hurufnya kecil menggunakan fungsi rekursif } */
	if n == 0 {
		return
	}

	lowercase(s, n-1)

	fmt.Print(string(s[n-1] + 32))
}

func hitungKapital(kata string, index int) int {
	if index == len(kata) {
		return 0
	}

	if 'A' <= kata[index] && kata[index] <= 'Z' {
		return 1 + hitungKapital(kata, index+1)
	}

	return hitungKapital(kata, index+1)
}

func UnAritmatika(a, b, n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	F.S. Fungsi mengembalikan hasil aritmatika }*/
	if n == 1 {
		return a
	} else {
		return b + UnAritmatika(a, b, n-1) // gunakan baris ini untuk memproses perhitungan aritmatika rekursif
	}
}

func hitungLowercase(kata string, index int) int {
	/*{ I.S. Terdefinisi parameter kata sebagai input string
	F.S. menampilkan jummlah  jumlah huruf lowercase dari kata menggunakan fungsi rekursif*/
	if index == len(kata) { // masukan kondisi untuk mengecek index dengan jumlah string
		return 0
	}
	if 'a' <= kata[index] && kata[index] <= 'z' {
		return 1 + hitungLowercase(kata, index+1) // gunakan fungsi rekursif yang tepat
	}
	return hitungLowercase(kata, index+1)
}

func sumCubic(n int) int {
	/*{ I.S. Terdefinisi sebuah bilangan bulat positif n.
	F.S. Fungsi mengembalikan hasil perpangkatan 3 */
	if n == 1 {
		return 1
	} else {
		return n*n*n + sumCubic(n-1) // gunakan baris ini untuk memproses perhitungan rekursif
	}
}

func hitungSKS(totalSKS int) int {
	var kodeMatkul string
	fmt.Scan(&kodeMatkul)

	if kodeMatkul == "0" {
		return totalSKS
	}

	sks := 0

	if kodeMatkul[0] == 'A' || kodeMatkul[0] == 'B' || kodeMatkul[0] == 'C' {
		sks = 2
	} else if kodeMatkul[0] == 'D' || kodeMatkul[0] == 'E' || kodeMatkul[0] == 'F' {
		sks = 3
	}

	return hitungSKS(totalSKS + sks)
}

func main() {
	var totalSKS, sks int
	sks = 0
	totalSKS = hitungSKS(sks)
	fmt.Print(totalSKS)
}
