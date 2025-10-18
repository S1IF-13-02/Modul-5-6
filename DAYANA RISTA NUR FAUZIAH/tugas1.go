package main

import "fmt"

func main() {

	var n int
	var jumlah int

	fmt.Print("Masukkan bilangan bulat positif (n): ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {

		jumlah = jumlah + i

	}

	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, jumlah)
}
