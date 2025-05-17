package main

import (
	"fmt"
)

type Mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type tabMahasiswa [1000]Mahasiswa

func inputData(m *tabMahasiswa, n *int) {
	var nama, nim, kelas, jurusan string
	var ipk float64
	fmt.Print("Masukkan jumlah data: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i+1)
		fmt.Scan(&nama)
		fmt.Printf("Masukkan nim mahasiswa ke-%d: ", i+1)
		fmt.Scan(&nim)
		fmt.Printf("Masukkan kelas mahasiswa ke-%d: ", i+1)
		fmt.Scan(&kelas)
		fmt.Printf("Masukkan jurusan mahasiswa ke-%d: ", i+1)
		fmt.Scan(&jurusan)
		fmt.Printf("Masukkan IPK mahasiswa ke-%d: ", i+1)
		fmt.Scan(&ipk)
		fmt.Println()

		m[i] = Mahasiswa{
			nama:    nama,
			nim:     nim,
			kelas:   kelas,
			jurusan: jurusan,
			ipk:     ipk,
		}
	}
}

// Sorting data berdasarkan NIM (ascending)
func sortByNIM(m *tabMahasiswa, n int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if m[i].nim > m[j].nim {
				m[i], m[j] = m[j], m[i]
			}
		}
	}
}

func cetakMahasiswa(m tabMahasiswa, n int) {
	fmt.Println("Daftar Mahasiswa:")
	for i := 0; i < n; i++ {
		fmt.Printf("%-10s %-20s %-15s %-15s %.2f\n", m[i].nama, m[i].nim, m[i].kelas, m[i].jurusan, m[i].ipk)
	}
}

func binarySearchNIM(m tabMahasiswa, n int, target string) int {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if m[mid].nim == target {
			return mid
		} else if m[mid].nim < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	var data tabMahasiswa
	var jmlhData int
	var target string
	inputData(&data, &jmlhData)
	sortByNIM(&data, jmlhData)
	cetakMahasiswa(data, jmlhData)
	fmt.Println()
	fmt.Print("Masukkan NIM yang ingin dicari: ")
	fmt.Scan(&target)
	idx := binarySearchNIM(data, jmlhData, target)
	if idx != -1 {
		fmt.Println("Data ditemukan:")
		fmt.Print(data[idx].nama)
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}
