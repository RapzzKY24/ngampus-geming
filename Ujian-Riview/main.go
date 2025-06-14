package main

import "fmt"

const NMAX = 10

type Snack struct {
	id    string
	nama  string
	harga int
	stok  int
}

type tabSnack [NMAX]Snack

func bacaData(t *tabSnack, n *int) {
	var inputId, inputNama string
	var inputHarga, inputStok int
	for *n < NMAX {
		fmt.Println("Masukkan data snack: ")
		fmt.Print("Masukkan id snack: ")
		fmt.Scan(&inputId)
		fmt.Print("Masukkan data nama snack: ")
		fmt.Scan(&inputNama)
		fmt.Print("Masukkan data harga snack: ")
		fmt.Scan(&inputHarga)
		fmt.Print("Masukkan data stok snack: ")
		fmt.Scan(&inputStok)
		if inputId == "none" && inputNama == "none" && inputHarga == 0 && inputStok == 0 {
			break
		}
		t[*n] = Snack{
			id:    inputId,
			nama:  inputNama,
			harga: inputHarga,
			stok:  inputStok,
		}
		(*n)++
	}
}

func cetak(t tabSnack, n int) {
	fmt.Println("===============================================")
	fmt.Printf("| %-10s | %-15s | %-7s | %-5s |\n", "ID", "Nama", "Harga", "Stok")
	fmt.Println("===============================================")
	for i := 0; i < n; i++ {
		fmt.Printf("| %-10s | %-15s | %-7d | %-5d |\n", t[i].id, t[i].nama, t[i].harga, t[i].stok)
	}
	fmt.Println("===============================================")
}

func namaSnackSeq(t tabSnack, n int, x string) string {
	for i := 0; i < n; i++ {
		if t[i].id == x {
			return t[i].nama
		}
	}
	return "none"
}

func omset(t tabSnack, n int) int {
	omset := 0
	for i := 0; i < n; i++ {
		omset += t[i].harga * t[i].stok
	}
	return omset
}

func urutIdMenarik(t *tabSnack, n int) {
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if (*t)[j].id < (*t)[min].id {
				min = j
			}
		}
		if min != i {
			(*t)[i], (*t)[min] = (*t)[min], (*t)[i]
		}
	}
	cetak(*t,n)
}

func urutHargaMenurun(t *tabSnack, n int) {
	for i := 1; i < n; i++ {
		key := (*t)[i]
		j := i - 1
		for j >= 0 && (*t)[j].harga < key.harga {
			(*t)[j+1] = (*t)[j]
			j--
		}
		(*t)[j+1] = key
	}
	cetak(*t,n)
}

func namaSnackBin(t *tabSnack, n int, x int) string {
	left := 0
	right := n - 1
	for left <= right {
		mid := (left + right) / 2
		if (*t)[mid].harga == x {
			return (*t)[mid].nama
		} else if (*t)[mid].harga < x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return "barang tidak ditemukan"
}


func main() {
	var data tabSnack
	var jmlhData int
	var keyword int
	var id string

	fmt.Println("===============================================")
	fmt.Println("          DATA SNACK      ")
	fmt.Println("===============================================")

	bacaData(&data, &jmlhData)

	fmt.Println("\nDaftar Snack yang Tersedia:")
	cetak(data, jmlhData)

	fmt.Print("\nMasukkan ID snack yang ingin dicari: ")
	fmt.Scan(&id)
	nama := namaSnackSeq(data, jmlhData, id)
	if nama != "none" {
		fmt.Println("Nama Snack dengan ID", id, ":", nama)
	} else {
		fmt.Println("Snack dengan ID", id, "tidak ditemukan.")
	}

	fmt.Printf("\nTotal Omset Semua Snack: Rp %d\n", omset(data, jmlhData))

	fmt.Println("\nData snack dengan ID urut ascending:")
	urutIdMenarik(&data, jmlhData)

	fmt.Println("\nData snack dengan harga urut descending:")
	urutHargaMenurun(&data, jmlhData)

	fmt.Print("\nMasukkan harga snack yang ingin dicari: ")
	fmt.Scan(&keyword)
	hasil := namaSnackBin(&data, jmlhData, keyword)
	if hasil != "barang tidak ditemukan" {
		fmt.Println("Snack dengan harga", keyword, "adalah:", hasil)
	} else {
		fmt.Println("Snack dengan harga", keyword, "tidak ditemukan.")
	}
}
