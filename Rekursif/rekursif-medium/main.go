package main

import "fmt"

func jumlahDeretGenap(x int) int {
	if x == 0 {
		return 0
	} else if x%2 != 0 {
		return jumlahDeretGenap(x - 1)
	}
	return x + jumlahDeretGenap(x-2)
}

func multiple(M, N int) int {
	if M == 0 || N == 0 {
		return 0
	}
	return M + multiple(M, N-1)
}

func pangkat(M, N int) int {
	if N <= 0 {
		return 1
	}
	return M * pangkat(M, N-1)
}

func bagi(m, n int) int {
	if m < n {
		return 0
	} else {
		return 1 + bagi(m-n, n)
	}
}

func tabungan(t, n int) int {
	if n == 1 {
		return t // Basis rekursi: minggu pertama, langsung return saldo awal
	}
	return tabungan(t+(n-1)*2500, n-1) // Tambah (n-1)*2500 setiap minggu
}

func jumlahtiapdigit(n int) int {
	/*{Terdefinisi sebuah bilangan bulat n.
	  Fungsi mengembalikan hasil penjumlahan dari tiap digit n. }

	  hint :  gunakan modulo dan penjumlahan dengan fungsi rekursif*/
	if n == 0 {
		return 0 // Basis rekursi: n = 0, langsung return 0
	}
	return n%10 + jumlahtiapdigit(n/10) // Ambil digit terakhir, tambahkan dengan hasil rekursi
}

func maksimum(max int) int {
	var x int
	fmt.Scan(&x)
	if x == 0 {
		return max // Basis rekursi: jika input 0, kembalikan nilai maksimum
	}
	if x > max {
		max = x // Perbarui nilai maksimum jika input lebih besar
	}
	return maksimum(max) // Rekursi dengan nilai maksimum yang diperbarui
}

func decimalToQuarternary(n int) int {
	/*I.S. Terdefinisi nilai bilangan bulat n.
	  F.S. fungsi mengembalikan nilai 0 jika n = 0, jika nilai n > 0 maka akan mengembalikan nilai Quarternary*/
	if n == 0 {
		return 0
	} else {
		// Konversi desimal ke quarternary menggunakan modulo 4 dan rekursi
		return (n % 4) + 10*decimalToQuarternary(n/4)
	}
}

func hitungJam(a int) int {
	/*I.S. Terdefinisi nilai bilangan bulat a.
	F.S. fungsi mengembalikan jumlah jam bernilai integer*/
	if a == 2 {
		return 0 // Basis rekursi: jika jumlah bakteri adalah 2, maka jam = 0
	}

	return 1 + hitungJam(a/2) // Setiap kali bakteri membelah, tambahkan 1 jam
}

func banyakBilanganGenap(N int) int {
	/*  fungsi mengembalikan sebuah bilangan bulat yang menyatakan
	banyaknya bilangan genap dari 2 hingga N */
	if N < 2 {
		return 0 // Basis rekursi: jika N kurang dari 2, tidak ada bilangan genap
	} else if N%2 != 0 {
		return banyakBilanganGenap(N - 1) // Lewati bilangan ganjil
	}
	return 1 + banyakBilanganGenap(N-2) // Hitung bilangan genap dan lanjutkan rekursi
}

func main() {
	fmt.Println(banyakBilanganGenap(1))
	fmt.Println(banyakBilanganGenap(2))
	fmt.Println(banyakBilanganGenap(4))
	fmt.Println(banyakBilanganGenap(5))
}
