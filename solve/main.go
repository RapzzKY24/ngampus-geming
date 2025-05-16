package main

import (
	"fmt"
	"math"
)

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
	var radius, massaKanan, massaKiri, tinggiKiri, tinggiKanan float64

	fmt.Print("Masukkan radius tabung: ")
	fmt.Scanln(&radius)

	fmt.Print("Masukkan massa tabung kiri: ")
	fmt.Scanln(&massaKiri)

	fmt.Print("Masukkan tinggi tabung kiri: ")
	fmt.Scanln(&tinggiKiri)

	fmt.Print("Masukkan massa tabung kanan: ")
	fmt.Scanln(&massaKanan)

	fmt.Print("Masukkan tinggi tabung kanan: ")
	fmt.Scanln(&tinggiKanan)

	display(radius, massaKanan, massaKiri, tinggiKiri, tinggiKanan)
}
