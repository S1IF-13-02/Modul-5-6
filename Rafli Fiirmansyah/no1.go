package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah %d", n, sum)
}
