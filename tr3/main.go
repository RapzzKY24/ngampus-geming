package main

import "fmt"

func pesanEror(flag int, m int) {
	message := ""
	if flag == 0 {
		message = "error"
	} else if flag == 1 {
		message = "warning"
	} else if flag == 2 {
		message = "Informasi"
	}
	fmt.Println(m, message)
}

func main() {
	var flag, m int
	fmt.Println("Masukkan bilangan dan pesan: ")
	fmt.Scan(&flag, &m)
	pesanEror(flag, m)
}
