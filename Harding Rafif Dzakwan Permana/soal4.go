package main

import "fmt"

func main() {
	var n int
	var hasil int = 1

	// Input bilangan
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	// Perulangan untuk menghitung faktorial
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}

	// Output hasil faktorial
	fmt.Println("Hasil faktorial:", hasil)
}
