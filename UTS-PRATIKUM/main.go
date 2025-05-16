package main

import "fmt"

func biayaKirim(beratPaket float64, jenisLayanan string) float64 {
	tarif := map[string]float64{"reguler": 10000, "kilat": 15000}
	totalBiaya := beratPaket * tarif[jenisLayanan]

	if beratPaket > 10 {
		totalBiaya *= 0.8
	}

	return totalBiaya
}

func hitungMasaHidup(daya *float64, jam *int) {
	for *daya >= 50 {
		if *daya > 500 {
			*daya *= 0.9
		} else {
			*daya *= 0.8
		}
		(*jam)++
	}
	fmt.Printf("Kristal bertahan selama %d jam sebelum mati. Daya kristal saat ini: %.2f\n", *jam, *daya)
}

func s(n int, k int) int {
	if n < k {
		return 0
	} else {
		return n + s(n-k, k)
	}
}

func g(i int) int {
	if i == 0 {
		return 0
	} else if i > 5 {
		return 500 + g(i-1)
	} else {
		return 200 + g(i-1)
	}
}

func k(n, m int) int {
	return (n * m) - g(n)
}

//g(0) = 0
// g(1) = 200 + g(0) = 200
// g(2) = 200 + 200 = 400
// g(3) = 200 + 400 = 600
// g(4) = 200 + 600 = 800
// g(5) = 200 + 800 = 1000
// g(6) = 500 + 1000 = 1500

// k(6, 10000) = (6 * 10000) - g(6)
//             = 60000 - 1500
//             = 58500

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Println(k(6, 10000))
}
