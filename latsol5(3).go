package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukkan angka perulangan: ")
	fmt.Scan(&a, &b)
	hasil := 0
	for c = 0; c < b; c++ {
		hasil = hasil + a
	}
	fmt.Println("Hasil Perkalian: ", hasil)
}
