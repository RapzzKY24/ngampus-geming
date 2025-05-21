package main

import "fmt"

const NMAX int = 2022

type Student struct {
	nama, sid string
	GPA       float64
}
type tabMhs [NMAX - 1]Student

func inputDataMhs(t *tabMhs, n *int) {
	var inputNama, inputSid string
	var inputGpa float64
	fmt.Print("Masukkan total mahasiswa: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan nama mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputNama)
		fmt.Printf("Masukkan student id mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputSid)
		fmt.Printf("Masukkan GPA mahasiswa ke-%d: ", i+1)
		fmt.Scan(&inputGpa)
		(*t)[i] = Student{
			nama: inputNama,
			sid:  inputSid,
			GPA:  inputGpa,
		}
		fmt.Println()
	}
}

func sortingAsc(t *tabMhs, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if t[j].GPA < t[minIdx].GPA {
				minIdx = j
			}
		}
		if minIdx != i {
			(*t)[i], (*t)[minIdx] = (*t)[minIdx], (*t)[i]
		}
	}
}

func showArray(t tabMhs, n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i+1, t[i].nama, t[i].sid, t[i].GPA)
	}
}

func main() {
	var dataMhs tabMhs
	var jmlhDataMhs int
	inputDataMhs(&dataMhs, &jmlhDataMhs)
	sortingAsc(&dataMhs, jmlhDataMhs)
	showArray(dataMhs, jmlhDataMhs)
}
