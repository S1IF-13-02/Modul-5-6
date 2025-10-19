package main

import "fmt"

func main() {
	var x, y, hasil int

	fmt.Print("Masukkan bilangan yang akan dipangkatkan: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&y)

	hasil = 1

	for i := 1; i <= y; i++{
		hasil = hasil * x
	}
	fmt.Print("Hasil perpangkatan: ", hasil)

}