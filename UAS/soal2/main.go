package main

import "fmt"

type Dosen struct {
	nama, nip string
	gaji      int
}

const MAXDATA int = 1000

type tabDosen [MAXDATA]Dosen

func inputDataDosen(t *tabDosen, n *int) {
	var inputNama, inputNIP string
	var inputGaji int
	fmt.Print("Masukkan jumlah dosen yang ingin diinput : ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukkan nama dosen ke-%d :  ",i+1)
		fmt.Scan(&inputNama)
		fmt.Printf("Masukkan NIP dosen ke-%d : ",i+1)
		fmt.Scan(&inputNIP)
		fmt.Printf("Masukkan Gaji dosen ke-%d : ",i+1)
		fmt.Scan(&inputGaji)
		fmt.Println()
		(*t)[i] = Dosen{
			nama: inputNama,
			nip: inputNIP,
			gaji: inputGaji,
		}
	}
}

func selectionSort(t *tabDosen,n int){
	for i := 0 ; i < n-1 ; i++{
		max := i
		for j := i +1 ; j < n ; j++{
			if (*t)[j].gaji > (*t)[max].gaji{
				max = j
			}
		}
		if max != i {
			 (*t)[i],(*t)[max] = (*t)[max],(*t)[i]
		}
	}
}



// binery search
func cariDosen (t *tabDosen, n int,gaji int)int{
	selectionSort(t,n)
	left := 0
	right := n-1
	for left <= right {
		mid := (left+right)/2
		if (*t)[mid].gaji == gaji{
			return mid
		}else if(*t)[mid].gaji < gaji{
			right = mid -1
		}else {
			left = mid +1
		}
	}
	return -1
}
func printDosen(t tabDosen, n int, gaji int) {
	binarySearchResult := cariDosen(&t, n, gaji)
	if binarySearchResult >= 0 {
		fmt.Println("Data dosen ditemukan:")
		fmt.Printf("Nama: %s, NIP: %s, Gaji: %d\n", t[binarySearchResult].nama, t[binarySearchResult].nip, t[binarySearchResult].gaji)
	} else {
		fmt.Printf("Dosen dengan gaji %d tidak ditemukan\n", gaji)
	}
}

func main(){
	var dataDosen tabDosen
	var jmlhDosen int
	var gajiTarget int
	inputDataDosen(&dataDosen,&jmlhDosen)
	fmt.Print("Masukkan gaji dosen yang ingin dicari: ")
	fmt.Scan(&gajiTarget)

	fmt.Println() // biar ada spasi

	printDosen(dataDosen,jmlhDosen,gajiTarget)
}
