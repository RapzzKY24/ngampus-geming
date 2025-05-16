package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	titikPusat Titik
	radius     int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.titikPusat, p) <= float64(c.radius)
}

func main() {
	var x1, y1, r1, x2, y2, r2, px, py int

	fmt.Println("Masukkan koordinat x dan y untuk titik pusat lingkaran 1:")
	fmt.Scan(&x1, &y1)
	fmt.Println("Masukkan radius lingkaran 1:")
	fmt.Scan(&r1)
	c1 := Lingkaran{Titik{x1, y1}, r1}

	fmt.Println("Masukkan koordinat x dan y untuk titik pusat lingkaran 2:")
	fmt.Scan(&x2, &y2)
	fmt.Println("Masukkan radius lingkaran 2:")
	fmt.Scan(&r2)
	c2 := Lingkaran{Titik{x2, y2}, r2}

	fmt.Println("Masukkan koordinat x dan y untuk titik sembarang:")
	fmt.Scan(&px, &py)
	p := Titik{px, py}

	inC1 := didalam(c1, p)
	inC2 := didalam(c2, p)

	if inC1 && inC2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inC1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inC2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
