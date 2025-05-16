package main

import "fmt"

func temperatureConverter(temperature float64) (float64, float64, float64) {
	reamur := temperature * 4 / 5
	farenheit := (temperature * 9 / 5) + 32
	kelvin := temperature + 273.15
	return reamur, farenheit, kelvin
}

func pecahUang(uang int, k10 *int, k2 *int, k1 *int) {
	*k10 = uang / 10000
	sisa := uang % 10000
	*k2 = sisa / 2000
	sisa = sisa % 2000
	*k1 = sisa / 1000
}

func mengurutkan(x *int, y *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
}

func main() {
	var temperature float64
	fmt.Print("Masukkan suhu dalam bentuk celcius: ")
	fmt.Scan(&temperature)
	fmt.Println(temperatureConverter(temperature))

}
