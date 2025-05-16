package main

import "fmt"

const NMAX = 1000

type tabMhs [NMAX]int

func inputData(t *tabMhs, n int) {
	var input int
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		t[i] = input
	}
}

func SortigGPA(t tabMhs, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := 1 + i; j < n; j++ {
			if t[j] < t[minIdx] {
				minIdx = j
			}
		}
		t[i], t[minIdx] = t[minIdx], t[i]
	}
	fmt.Println(t[:n])
}

func main() {
	var data tabMhs
	inputData(&data, 5)
	SortigGPA(data, 5)
}
