package main

import "fmt"

type Student struct {
	name, sid string
	age       int
	gpa       float64
}

const NMAX = 2025

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
		s[i] = Student{
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
		key := (*s)[i]
		j := i - 1
		for j >= 0 && (*s)[j].age < key.age {
			(*s)[j+1] = (*s)[j]
			j--
		}
		(*s)[j+1] = key
	}
}

func main() {
	var dataMhs tabStudent
	var jmlhMhs int
	inputStudent(&dataMhs,&jmlhMhs)
	selectionSortGpaASC(&dataMhs, jmlhMhs)
	fmt.Println("Data mahasiswa setelah diurutkan berdasarkan GPA (Ascending):")
	for i := 0; i < jmlhMhs; i++ {
		fmt.Printf("Nama: %s, SID: %s, Umur: %d, GPA: %.2f\n", dataMhs[i].name, dataMhs[i].sid, dataMhs[i].age, dataMhs[i].gpa)
	}
	insertionSortAgeDesc(&dataMhs, jmlhMhs)
	fmt.Println("\nData mahasiswa setelah diurutkan berdasarkan Umur (Descending):")
	for i := 0; i < jmlhMhs; i++ {
		fmt.Printf("Nama: %s, SID: %s, Umur: %d, GPA: %.2f\n", dataMhs[i].name, dataMhs[i].sid, dataMhs[i].age, dataMhs[i].gpa)
	}
}
