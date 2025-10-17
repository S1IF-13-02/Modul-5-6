package main

import "fmt"

func main() {
	var n, j, hasil int

	fmt.Print("masukan penjumlahan sampai n : ")
	fmt.Scan(&n)
	for j = 1; j <= n; j++ {

		hasil = hasil + j

	}
	fmt.Print(hasil)

}
