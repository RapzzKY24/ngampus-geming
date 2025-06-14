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
		t[i] = Dosen{
			nama: inputNama,
			nip: inputNIP,
			gaji: inputGaji,
		}
	}
}

//sequental search
func cariDosen(t tabDosen,n int,nip string)int{
	for i := 0 ; i < n ; i++{
		if t[i].nip == nip{
			return i
		}
	}
	return -1
}

func printDosen(t tabDosen, n int,nip string){
	seqSearchResult := cariDosen(t,n,nip)
	if seqSearchResult >= 0 {
		fmt.Println("Dosen ditemukan: ",t[seqSearchResult].nama,t[seqSearchResult].nip,t[seqSearchResult].gaji)
	}else{
		fmt.Printf("Dosen dengan NIP %s tidak ditemukan",nip)
	}
}

func main(){
	var dataDosen tabDosen
	var jmlhDosen int
	var nipTarget string
	inputDataDosen(&dataDosen,&jmlhDosen)
	fmt.Print("Masukkan NIP dosen yang ingin dicari: ")
	fmt.Scan(&nipTarget)

	fmt.Println() // biar ada spasi

	printDosen(dataDosen,jmlhDosen,nipTarget)
}
