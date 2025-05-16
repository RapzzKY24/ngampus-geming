package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	Nama    string
	NIM     int
	Jurusan string
}

func SequentialSearch(mahasiswaList []Mahasiswa, keyword string) *Mahasiswa {
	for _, m := range mahasiswaList {
		if m.Jurusan == keyword {
			return &m
		}
	}
	return nil
}

func main() {
	mahasiswaList := []Mahasiswa{
		{
			Nama:    "Payme Risky",
			NIM:     12345678,
			Jurusan: "Teknik Informatika",
		},
		{
			Nama:    "Muhammad Teguh",
			NIM:     87654321,
			Jurusan: "Sistem Informasi",
		},
		{
			Nama:    "Lexa Namiko",
			NIM:     11223344,
			Jurusan: "Teknik Elektro",
		},
		{
			Nama:    "Sulthan Andyno",
			NIM:     44332211,
			Jurusan: "Teknik Mesin",
		},
		{
			Nama:    "Rasya Naufal",
			NIM:     55667788,
			Jurusan: "Teknik Sipil",
		},
		{
			Nama:    "LeoKhansa Naghara",
			NIM:     55667788,
			Jurusan: "Teknik Sipil",
		},
		{
			Nama:    "Ruben Febrian",
			NIM:     31242441,
			Jurusan: "Teknik Industri",
		},
		{
			Nama:    "Sulthan Balya",
			NIM:     432424242,
			Jurusan: "Blok M Bca",
		},
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama : ")
	keyword, _ := reader.ReadString('\n')
	keyword = strings.TrimSpace(keyword)
	result := SequentialSearch(mahasiswaList, keyword)
	if result != nil {
		fmt.Printf("Data ditemukan:\nNama: %s\nNIM: %d\nJurusan: %s\n", result.Nama, result.NIM, result.Jurusan)
	} else {
		fmt.Println("Data tidak ditemukan")
	}

}
