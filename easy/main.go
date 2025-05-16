package main

import "fmt"

func penjumlahanInterval(M int, N int) {
	/*
	   I.S. terdefinisi dua buah bilangan bulat M dan N
	   F.S. menampilkan hasil penjumlahan dari M hingga N
	*/
	count := 0
	for i := M; i <= N; i++ {
		count += i
	}
	fmt.Println(count)
}

func faktor(a, b int) {
	/*
	   I.S. berupa dua buah bilangan bulat positif, a dan b.
	   F.S. tercetak "Ya" jika a adalah faktor dari b atau "Tidak" jika bukan.
	*/
	if b%a == 0 {
		fmt.Printf("Ya, %d faktor %d\n", a, b)
	} else {
		fmt.Printf("Tidak, %d bukan faktor %d\n", a, b)
	}
}

// Buatlah prosedur saja
func cetakHitungJumlahRata(n, m int, jum *int, rata *float64) {
	/* I.S.: Terdefinisi bilangan bulat n dan m
	   F.S.: jum berisi nilai penjumlahan bilangan bulat dari n hingga m
			 dan rata berisi nilai rata-rata bilangan bulat dari n hingga m */
	*jum = 0
	totalBilangan := m - n + 1 // Menghitung jumlah bilangan dari n hingga m
	for i := n; i <= m; i++ {
		fmt.Println(i)
		*jum += i
	}
	*rata = float64(*jum) / float64(totalBilangan) // Menghitung rata-rata
}

func goNoGo(hujanTurun string, bawaPayung string) {
	/*
	   I.S. terdefinisi dua buah string hujanTurun dan bawaPayung ("ya" atau "tidak").
	   F.S. menampilkan string "berangkat" bila hujan tidak turun atau membawa payung,
	        atau menampilkan "diam di rumah" bila sebaliknya.
	*/

	if hujanTurun == "tidak" {
		fmt.Println("berangkat")
	} else if hujanTurun == "ya" {
		if bawaPayung == "ya" {
			fmt.Println("berangkat")
		} else {
			fmt.Println("diam di rumah")
		}
	}
}

func cetak(n int, x string) {
	/*
	   I.S. terdefinisi sebuah bilangan bulat, n, yang menyatakan banyaknya pencetakan dan sebuah string, x, yang akan dicetak.
	   F.S. tercetak string x sebanyak n kali.
	*/
	for i := 1; i <= n; i++ {
		fmt.Println(x)
	}

}
