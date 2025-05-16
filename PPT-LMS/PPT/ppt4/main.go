package main

import (
	"fmt"
	"math"
)

func zoomIn(width, height, scala int) {
	panjangZoomIn, lebarZoomIn := width*scala, height*scala
	fmt.Println(panjangZoomIn, lebarZoomIn)
}

func zoomOut(width, height, scala int) {
	panjangZoomOut, lebarZoomOut := width/scala, height/scala
	fmt.Println(panjangZoomOut, lebarZoomOut)
}

func volume(radius, tinggi float64) float64 {
	return 3.14 * radius * radius * tinggi
}

func hitungMassa(radius, massa, tinggi float64) float64 {
	return volume(radius, tinggi) * massa
}

func display(radius, massaKanan, massaKiri, tinggiKiri, tinggiKanan float64) {
	massaTabungKiri := hitungMassa(radius, tinggiKiri, massaKiri)
	massaTabungKanan := hitungMassa(radius, tinggiKanan, massaKanan)

	if massaTabungKanan == massaTabungKiri {
		fmt.Println("Balance")
	} else {
		fmt.Println(math.Abs(massaTabungKanan - massaTabungKiri))
	}
}

func main() {
	var width, height, scala int
	var choice string

	fmt.Print("Masukkan lebar gambar: ")
	fmt.Scan(&width)
	fmt.Print("Masukkan tinggi gambar: ")
	fmt.Scan(&height)
	fmt.Print("Masukkan skala: ")
	fmt.Scan(&scala)

	fmt.Print("Pilih mode (zoomIn/zoomOut): ")
	fmt.Scan(&choice)

	if choice == "zoomIn" {
		zoomIn(width, height, scala)
	} else if choice == "zoomOut" {
		zoomOut(width, height, scala)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
}
