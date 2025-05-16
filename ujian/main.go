package main

import "fmt"

func jumlahPerkalian(node1, node2, bobot1, bobot2 int) int {
	perkalianNode1 := node1 * bobot1
	perkalianNode2 := node2 * bobot2
	result := perkalianNode1 + perkalianNode2
	return result
}

func aktivitas(x int) int {
	if x < 0 {
		return 0
	} else {
		return x
	}
}

func geometriRekursif(n float64) float64 {
	if n == 1 {
		return 1
	} else {
		return 0.5 * geometriRekursif(n-1)
	}
}

func main() {
	var input1, input2, b1, b2, b3, b4, b5, b6 int
	fmt.Scan(&input1, &input2)
	fmt.Scan(&b1, &b2, &b3, &b4, &b5, &b6)
	node1 := jumlahPerkalian(input1, input2, b1, b2)
	nodeHiden1 := aktivitas(node1)
	node2 := jumlahPerkalian(input1, input2, b3, b4)
	nodeHiden2 := aktivitas(node2)
	result := nodeHiden1*b5 + nodeHiden2*b6
	fmt.Println(result)

	var n float64
	fmt.Print("Masukkan bilangan n :")
	fmt.Scan(&n)
	fmt.Println(geometriRekursif(n))
}
