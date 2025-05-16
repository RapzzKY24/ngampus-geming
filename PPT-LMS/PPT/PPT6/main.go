package main

import "fmt"

type Rectange struct {
	length, width int
	color         string
	property      Geometry
}

type Geometry struct {
	area      int
	perimeter int
}

func isiData(persegi *Rectange) {
	fmt.Print("Masukkan lebar persegi panjang: ")
	fmt.Scan(&persegi.length)
	fmt.Print("Masukkan panjang persegi panjang: ")
	fmt.Scan(&persegi.width)
	fmt.Print("Masuukan warna persegi panjang: ")
	fmt.Scan(&persegi.color)
}

func calculateArea(persegi Rectange) int {
	return persegi.length * persegi.width
}

func calculatePerimeter(persegi Rectange) int {
	return 2 * (persegi.length + persegi.width)
}

func hitung(persegi *Rectange) {
	persegi.property.perimeter = calculatePerimeter(*persegi)
	persegi.property.area = calculateArea(*persegi)
	fmt.Println("Luas dari persegi panjang tersebut adalah:", persegi.property.area)
	fmt.Println("Keliling dari persegi panjang tersebut adalah:", persegi.property.perimeter)
}

func main() {
	var dataPersegi Rectange
	isiData(&dataPersegi)
	hitung(&dataPersegi)
}
