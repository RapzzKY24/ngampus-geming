package main

import "fmt"

const NMAX = 256

type tabArray [NMAX]int

func isiData(t *tabArray, n *int) {
	var inputData int
	*n = 0
	for *n < NMAX {
		fmt.Printf("Masukkan data ke array ke %d :", *n+1)
		fmt.Scan(&inputData)

		if inputData == 0 {
			break
		}
		t[*n] = inputData
		(*n)++
	}
}

func cetak(t tabArray, n int) []int {
	var array []int
	for i := 0; i < n; i++ {
		array = append(array, t[i])
	}
	return array
}
func reverse(t tabArray, n int) []int {
	array := cetak(t, n)
	reversedArray := make([]int, n)
	left := 0
	rigth := n - 1
	for left < n {
		reversedArray[left] = array[rigth]
		rigth--
		left++
	}
	return reversedArray
}

func palindrome(t tabArray, n int) string {
	array := cetak(t, n)
	reverse := reverse(t, n)
	isPalindrome := true
	for i := 0; i < n; i++ {
		if array[i] != reverse[i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		return "Palindrom"
	}
	return "Bukan Palindrom"
}

func main() {
	var jmlhData int
	var data tabArray
	isiData(&data, &jmlhData)
	fmt.Println(reverse(data, jmlhData))
	fmt.Println(palindrome(data, jmlhData))
}
