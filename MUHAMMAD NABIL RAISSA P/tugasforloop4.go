package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan masukan: ")
	fmt.Scan(&n)

	var hasil uint64 = 1

	for i := 1; i <= n; i++ {
		hasil = hasil * uint64(i)
	}

	fmt.Println(hasil)
}