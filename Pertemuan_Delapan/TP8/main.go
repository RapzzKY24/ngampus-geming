package main

import "fmt"

const NMAX = 1000

type tabIkan [NMAX]float64

func inputBerat(t *tabIkan, n *int, y *int) {
	fmt.Print("Masukkan jumlah ikan: ")
	fmt.Scan(n)

	fmt.Print("Masukkan kapasitas per wadah: ")
	fmt.Scan(y)

	for i := 0; i < *n && *n < NMAX; i++ {
		var berat float64
		fmt.Printf("Masukkan berat ikan ke-%d: ", i+1)
		fmt.Scan(&berat)
		t[i] = berat
	}
}

func pindahSesuaiKapasitas(t tabIkan, n int, y int) ([]float64, float64, float64) { // return 1 : total ikan di wadah , 2 : max , 3 : min
	var wadah []float64
	var currentWadah []float64
	totalBerat := 0.0

	for i := 0; i < n; i++ {
		currentWadah = append(currentWadah, t[i])
		totalBerat += t[i]

		if len(currentWadah) == y || i == n-1 {
			wadah = append(wadah, totalBerat)
			currentWadah = nil
			totalBerat = 0.0
		}
	}

	max := wadah[0]
	min := wadah[0]

	for i := 0; i < len(wadah); i++ {
		if wadah[i] > max {
			max = wadah[i]
		}

		if wadah[i] < min {
			min = wadah[i]
		}
	}

	return wadah, max, min
}

func rataRata(t tabIkan, n int, y int) float64 {
	jumlah := 0.0
	for i := 0; i < n; i++ {
		jumlah += t[i]
	}
	return float64(jumlah) / float64(y)
}

func main() {
	var t tabIkan
	var n, y int

	inputBerat(&t, &n, &y)

	wadah, max, min := pindahSesuaiKapasitas(t, n, y)

	fmt.Println(wadah)

	fmt.Printf("Rata-rata berat ikan: %.2f\n", rataRata(t, n, y))
	fmt.Printf("Berat minimum: %.2f\n", min)
	fmt.Printf("Berat maksimum: %.2f\n", max)
}
