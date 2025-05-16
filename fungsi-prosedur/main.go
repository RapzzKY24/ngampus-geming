package main

import (
	"fmt"
)

func kelipatanTiga(jumlah *int) {
	/* I.S. jumlah diinisialisasi dengan nilai 0
	   F.S. jumlah telah berisi jumlah dari bilangan-bilangan ganjil kelipatan 3
	*/
	count := 0
	var n int

	for {
		fmt.Scan(&n)
		if n < 0 {
			break
		}
		if n%3 == 0 && n%2 != 0 {
			count += n
		}
	}

	*jumlah = count
}

// Buatlah prosedur saja

func jumlahDeret(n, m int, jumlah *float64) {
	if n > m {
		return
	}
	/* I.S Terdefinisi dua bilangan bulat positif yaitu n dan m, dan variabel real jumlah bernilai nol
	   F.S Hasil dari jumlah deret n dan m ditampung oleh variabel jumlah*/
	*jumlah += 4.0 / float64(n)
	jumlahDeret(n+1, m, jumlah)
}

func audisi(j1, j2, j3 string) {
	/* I.S Terdefinisi 3 variabel string yang menyatakan vote dari ketiga juri, String "yes" berarti juri menyatakan lolos, sedangkan "no" menyatakan tidak lolos.
	   F.S Menampilkan string "lolos" atau "tidak lolos" bergantung dari nilai kondisinya.*/

	if j1 == "yes" && j2 == "yes" && j3 == "yes" {
		fmt.Println("lolos")
	} else {
		fmt.Println("tidak lolos")
	}
}

//buatlah procedure saja

func totalPelajar(a string, b int, totSD, totSMP, totSMA *int) {
	/* I.S. terdefisi pelajar dan jumlah
	   F.S. totSD adalah jumlah pelajar SD, totSMP adalah jumlah pelajar SMP, totSMA adalah jumlah pelajar SMA*/
	switch a {
	case "sd":
		*totSD += b
	case "smp":
		*totSMP += b
	case "sma":
		*totSMA += b
	}
}

func hitungSewaBus(k, p int) {
	/*
		I.S menerima dua bilangan bulat k (kapasitas bus) dan p (jumlah peserta).
		F.S menghitung dan mencetak jumlah bus yang diperlukan untuk mengangkut
		semua peserta dengan membagi jumlah peserta (p)
		dengan kapasitas bus (k), dan menambahkan satu jika ada sisa peserta yang
		tidak dapat diakomodasi penuh oleh bus.
	*/
	jumlahBus := p / k
	if p%k != 0 {
		jumlahBus++
	}
	fmt.Println(jumlahBus)
}

//buatlah procedure saja

func cetakBilangan(x, y, a, b int) {
	/* I.S. terdefinisi bilangan x dan y adalah bilangan bulat berdigit 3 buah dan a, b bilangan bulat 1 digit.
	   F.S. berupa bilangan bulat dari x hingga y, kecuali bilangan yang dimulai dengan bilangan a atau diakhiri bilangan b. */

	if x > y {
		return
	}

	if x/100 != a && x%10 != b {
		fmt.Println(x)
	}

	cetakBilangan(x+1, y, a, b)

}

// buatlah prosedurnya saja

func cetakTerbesar(b1, b2, b3, b4 int) {
	/*
		I.S menerima empat bilangan bulat b1, b2, b3, dan b4.
		F.S mencari dan mencetak bilangan terbesar di antara b1, b2, b3, dan b4.
	*/
	max := b1
	if b2 > max {
		max = b2
	}
	if b3 > max {
		max = b3
	}
	if b4 > max {
		max = b4
	}
	fmt.Println(max)
}

// buatlah prosedurnya saja

func skorTertinggiPanahan(s1, s2, s3 int, t1, t2, t3, maks *int) {
	/*
	   I.S menerima tiga skor panahan s1, s2, dan s3, serta tiga variabel penjumlahan skor t1, t2, dan t3,
	   dan pointer ke integer maks yang diinisialisasi dengan nilai awal 0.
	   F.S mengupdate total skor t1, t2, dan t3 dengan menambahkan s1, s2, dan s3 masing-masing. Selanjutnya, nilai
	   tertinggi dari tiga total skor (t1, t2, t3) ditentukan dan disimpan dalam variabel maks.
	*/
	*t1 += s1
	*t2 += s2
	*t3 += s3
	*maks = *t1
	if *t2 > *maks {
		*maks = *t2
	}
	if *t3 > *maks {
		*maks = *t3
	}
}

func skorDanPemenang(n int) {
	/*  I.S terdefinisi bilangan bulat n yang menyatakan banyaknya ronde
	F.S terdiri dari tiga nilai, yaitu skor petinju A, skor petinju B, dan karakter 'A', 'B', atau '0' */
	var skorA, skorB int

	for i := 0; i < n; i++ {
		var a1, a2, a3, b1, b2, b3 int
		fmt.Scan(&a1, &a2, &a3, &b1, &b2, &b3)
		skorA += a1 + a2 + a3
		skorB += b1 + b2 + b3
	}

	fmt.Printf("%d %d ", skorA, skorB)
	if skorA > skorB {
		fmt.Println("A")
	} else if skorB > skorA {
		fmt.Println("B")
	} else {
		fmt.Println("0")
	}
}

//Buatlah prosedur saja

func monetize(jumPelanggan int, jumJamTonton int) {
	/* I.S. Jumpelanggan dan jumJamTonton terderdefinisi
	   F.S. Menampilkan "sudah dapat dimonetisasi" atau "belum dapat dimonetisasi" sesuai dengan ketentuan soal*/
	if jumPelanggan >= 1000 && jumJamTonton >= 4000 {
		fmt.Println("sudah dapat dimonetisasi")
	} else {
		fmt.Println("belum dapat dimonetisasi")
	}

}
func sum(kar byte, jumlah *int) {
	/* I.S. Terdefinisi kar byte dan jumlah int
	   F.S. jumlah huruf selain 'a' dan 'i' terdefinisi di jumlah int */
	*jumlah = 0

	for kar != '.' {

		if kar != 'a' && kar != 'i' {
			*jumlah++
		}
		fmt.Scanf("%c", &kar)
	}
}
func telepati(a, b, c int) {
	/* I.S. Terdefinisi a b dan c sebagai bilangan bulan positif
	   F.S. menampilkan dua nilai yang menyatakan banyaknya pemain yang bersahabat dan tidak.
	*/

	var sahabat, notsahabat int

	sahabat = 0
	notsahabat = 0
	fmt.Scan(&a, &b, &c)

	for a != 0 || b != 0 || c != 0 { // lakukan kondisi pengulangan untuk memperoleh jumlah
		if a != b && b != c && a != c && (a+b == c || a+c == b || b+c == a) {
			sahabat++
		} else {
			notsahabat++
		}
		fmt.Scan(&a, &b, &c)
	}
	fmt.Println(sahabat, notsahabat) // tampilkan hasil 2 nilai yang menyatakan banyaknya pemain yang bersahabat dan tidak.
}

func suhu(temp int) {
	/* I.S. terdefinisi sebuah bilangan bulat yang menyatakan temperatur.
	   F.S. print string untuk menyatakan status cuaca berikut: "Freezing weather",
	   "Very Cold weather", "Cold weather", "Normal in Temp", "it's Hot", atau
	   "It's Very Hot". sesuai dengan list cuaca diatas
	*/
	switch {
	case temp < 0:
		fmt.Println("Freezing weather")
	case temp >= 0 && temp <= 9:
		fmt.Println("Very Cold weather")
	case temp >= 10 && temp <= 19:
		fmt.Println("Cold weather")
	case temp >= 20 && temp <= 29:
		fmt.Println("Normal in Temp")
	case temp >= 30 && temp <= 39:
		fmt.Println("It's Hot")
	default:
		fmt.Println("It's Very Hot")
	}
}

func main() {
	var jumlah int
	var kar byte
	fmt.Scanf("%c", &kar)
	sum(kar, &jumlah)
	fmt.Println(jumlah)
}
