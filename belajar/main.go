package main

import "fmt"

const bunga = 100

type tabBunga [bunga]string

func cari(t tabBunga, bunga string, n int) int {
	i := 0
	found := -1
	for i < n && found == -1 {
		if t[i] == bunga {
			found = 1
		}
		i++
	}
	return found
}

func delete(t tabBunga, n int, x string) {
	found := cari(t, x, n)
	if found == -1 {
		fmt.Println("Data tidak ditemukan")
	} else {
		i := found
		for i <= n-2 {
			t[i] = t[i+1]
			i = i + 1
		}
		n -= 1
	}
}

func main() {
	var t tabBunga
	var n int
	var choice int
	var x string

	n = 0

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Cari Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			if n < bunga {
				fmt.Print("Masukkan nama bunga: ")
				fmt.Scan(&x)
				t[n] = x
				n++
				fmt.Println("Data berhasil ditambahkan.")
			} else {
				fmt.Println("Data penuh, tidak bisa menambah data.")
			}
		case 2:
			fmt.Print("Masukkan nama bunga yang dicari: ")
			fmt.Scan(&x)
			found := cari(t, x, n)
			if found == -1 {
				fmt.Println("Data tidak ditemukan.")
			} else {
				fmt.Println("Data ditemukan:", t[found])
			}
		case 3:
			fmt.Print("Masukkan nama bunga yang ingin dihapus: ")
			fmt.Scan(&x)
			delete(t, n, x)
		case 4:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
