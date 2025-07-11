package main

import "fmt"

type Student struct {
	name, sid string
	age       int
	gpa       float64
}

const NMAX int = 2025

type tabStudent [NMAX]Student

func inputStudent(s *tabStudent, n *int) {
	var inputName, inputSID string
	var inputAge int
	var inputGPA float64
	fmt.Print("Berapa data mahasiswa yang mau diinput: ")
	fmt.Scan(n)
	for i := 0 ; i <* n ; i++{
		// fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputName)
		// fmt.Printf("Masukkan SID mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputSID)
		// fmt.Printf("Masukkan umur mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputAge)
		// fmt.Printf("Masukkan GPA mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputGPA)
		(*s)[i] = Student{
			name: inputName,
			sid:  inputSID,
			age:  inputAge,
			gpa:  inputGPA,
		}
	}
}

func selectionSortGpaASC(s *tabStudent,n int){
	for i := 0 ; i < n-1 ; i++{
		minIdx := i
		for j := i +1 ; j < n ; j++{
			if (*s)[j].gpa < (*s)[minIdx].gpa{
				minIdx = j
			}
		}
		if minIdx != i {
			(*s)[i],(*s)[minIdx] = (*s)[minIdx],(*s)[i]
		}
	}
}

func insertionSortAgeDesc(s *tabStudent, n int) {
	for i := 1; i < n; i++ {
		key := (*s)[i] // Ambil elemen yang akan disisipkan
		// Pindahkan elemen yang lebih kecil dari key ke satu posisi ke kanan
		j := i - 1 // Indeks elemen sebelumnya
		// Urutkan berdasarkan umur secara menurun
		for j >= 0 && (*s)[j].age < key.age {
			// Pindahkan elemen ke kanan
			(*s)[j+1] = (*s)[j]
			j--
		}
		// Sisipkan elemen key pada posisi yang benar
		(*s)[j+1] = key
	}
}

func cetakSorting(s *tabStudent, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Nama: %s, SID: %s, Umur: %d, GPA: %.2f\n", (*s)[i].name, (*s)[i].sid, (*s)[i].age, (*s)[i].gpa)
	}
}

func main() {
	var dataMhs tabStudent
	var jmlhMhs int
	inputStudent(&dataMhs,&jmlhMhs)
	selectionSortGpaASC(&dataMhs, jmlhMhs)
	fmt.Println("Data mahasiswa setelah diurutkan berdasarkan GPA (Ascending):")
	cetakSorting(&dataMhs, jmlhMhs)
	fmt.Println()
	insertionSortAgeDesc(&dataMhs, jmlhMhs)
	fmt.Println("Data mahasiswa setelah diurutkan berdasarkan Umur (Descending):")
	cetakSorting(&dataMhs, jmlhMhs)
}
