package main

import "fmt"

// Buatlah prosedur saja

func login(username, password string, totalGagalLogin *int) {
	/* I.S. terdefinisi string username dan password, dan bilangan bulat totalGagalLogin
	   F.S. totalGagalLogin berisi jumlah gagal login dari user */
	for username != "admin" || password != "admin" {
		*totalGagalLogin += 1
		fmt.Scan(&username, &password)
	}
}

//buatlah prosedur saja

func secMax(max, secondMax *int) {
	/* I.S. terdefinisi bilangan terbesar max, bilangan terbesar kedua secondMax.
	   dan inputan barisan bilangan bulat yang berakhir jika bilangan bulat sama dengan 0
	   F.S. max adalah bilangan bulat terbesar dan secondMax adalah bilangan bulat terbesar kedua */
	var input int
	*max, *secondMax = 0, 0
	for {
		fmt.Scan(&input)
		if input == 0 {
			break
		}
		if input > *max {
			*secondMax = *max
			*max = input
		} else if input > *secondMax {
			*secondMax = input
		}
	}
	fmt.Println(*secondMax)
}

//buat prosedur saja

func konsekutif(bilangan int) {
	s := fmt.Sprintf("%d", bilangan)

	if len(s) == 1 {
		fmt.Println(true)
		return
	}

	for i := 0; i < len(s)-1; i++ {
		digit1 := int(s[i] - '0')
		digit2 := int(s[i+1] - '0')

		if abs(digit1-digit2) != 1 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	konsekutif(bilangan)

}
