package main

import "fmt"

func SequentialSearch(fruits []string, target string) []int {
	var result []int
	for indeks, fruit := range fruits {
		if fruit == target {
			result = append(result, indeks)
		}
	}
	return result
}

// func SequentialSearch(data []int, target int) int {
// 	for i := 0; i < len(data); i++ {
// 		if data[i] == target {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func bubleSort(array []int) []int {
// 	for i := 0; i < len(array)-1; i++ {
// 		for j := 0; j < len(array)-1-i; j++ {
// 			if array[j] > array[j+1] {
// 				array[j], array[j+1] = array[j+1], array[j]
// 			}
// 		}
// 	}
// 	return array
// }

// func selectionSort(array []int) []int {
// 	for i := 0; i < len(array); i++ {
// 		minIndex := i
// 		for j := i + 1; j < len(array); j++ {
// 			if array[j] < array[minIndex] {
// 				minIndex = j
// 			}
// 		}
// 		array[i], array[minIndex] = array[minIndex], array[i]
// 	}
// 	return array
// }

// func maximum(array []int) []int {
// 	max := array[0]
// 	for i := 0; i < len(array); i++ {
// 		if array[i] > max {
// 			max = array[i]
// 		}
// 	}
// 	return []int{max}
// }

// func binerySearch(array []string, target string) []int {
// 	right := 0
// 	left := len(array) - 1
// 	for right <= left {
// 		mid := (left + right) / 2
// 		if array[mid] == target {
// 			return []int{mid}
// 		} else if array[mid] < target {
// 			right++
// 		} else {
// 			left--
// 		}
// 	}
// 	return []int{-1}
// }

func main() {
	var target string
	fruits := []string{"Apple", "Manggo", "Papaya", "Pineapple", "Orange", "Papaya"}
	fmt.Print("Masukkan target yang ingin dicari: ")
	fmt.Scan(&target)
	result := SequentialSearch(fruits, target)
	fmt.Println("Data berapa pada indeks : ", result)
}
