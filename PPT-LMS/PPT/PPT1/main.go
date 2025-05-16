package main

import "fmt"

func timeConversion(hours, minute, second int) int {
	return hours*3600 + minute*60 + second
}

func checkSegitiga(a, b, c int) bool {
	return a+b > c && a+c > b && c+b > a
}

func discount(total_Belanja float64, member bool) float64 {
	discount := 0.0
	if total_Belanja > 100000 {
		if member {
			discount = 0.1
		} else {
			discount = 0.05
		}
	}
	return total_Belanja * (1 - discount)
}

func snFibonannci(n int) int {
	total := 0
	u1, u2 := 0, 1
	for i := 0; i < n; i++ {
		total += u1
		u1, u2 = u2, u1+u2
	}
	return total
}

func faktorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * faktorial(x-1)
}

func permutasi(x, y int) (int, int, int) {
	var result int
	faktorialX := faktorial(x)
	faktorialY := faktorial(y)

	if x >= y {
		result = faktorialX / faktorial(x-y)
	} else if x < y {
		result = faktorialY / faktorial(y-x)
	}
	return faktorialX, faktorialY, result
}

func main() {
	// var bil int
	// fmt.Scan(&bil)
	// fmt.Println("Hasil dari total: ", snFibonannci(bil))
	// var hours, minute, second int
	// fmt.Scan(&hours, &minute, &second)
	// conversion := timeConversion(hours, minute, second)
	// fmt.Printf("Hasil konversi= %d\n", conversion)

	// var totalBelanja float64
	// var isMember int

	// fmt.Print("Masukkan total belanja: ")
	// fmt.Scan(&totalBelanja)

	// fmt.Print("Apakah Anda member? (1 = Ya, 0 = Tidak): ")
	// fmt.Scan(&isMember)

	// finalPrice := discount(totalBelanja, isMember == 1)

	// fmt.Printf("Total yang harus dibayar: Rp%.2f\n", finalPrice)

	// fmt.Println(snFibonannci(7))
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if checkSegitiga(a, b, c) {
		fmt.Println("Segitiga")
	} else {
		fmt.Println("Bukan segitiga")
	}

}
