package main

import (
	"fmt"
	"math"
)

func euclidean(x, cx, y, cy float64) float64 {
	distance := math.Pow(x-cx, 2) + math.Pow(y-cy, 2)
	return math.Sqrt(distance)
}

func posisiTitikLingkaran(r1, r2, x, y, cx1, cy1, cx2, cy2 float64) {
	lingkaranSatu := euclidean(x, cx1, y, cy1)
	lingkaranDua := euclidean(x, cx2, y, cy2)
	if lingkaranSatu <= r1 && lingkaranDua <= r2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	} else if lingkaranSatu <= r1 {
		fmt.Println("Titik berada di dalam lingkaran 1")
	} else if lingkaranDua <= r2 {
		fmt.Println("Titik berada di dalam lingkaran 2")
	} else {
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}
}

func main() {
	var r1, r2, x, y, cx1, cy1, cx2, cy2 float64
	fmt.Print("Masukkan koordinat lingkaran pertama: ")
	fmt.Scan(&cx1, &cy1)
	fmt.Print("Masukkan radius lingkaran pertama: ")
	fmt.Scan(&r1)
	fmt.Print("Masukkan koordinat lingkaran kedua: ")
	fmt.Scan(&cx2, &cy2)
	fmt.Print("Masukkan radius lingkaran kedua: ")
	fmt.Scan(&r2)
	fmt.Print("Masukkan koordinat titik sembarangan: ")
	fmt.Scan(&x, &y)
	posisiTitikLingkaran(r1, r2, x, y, cx1, cy1, cx2, cy2)
}
