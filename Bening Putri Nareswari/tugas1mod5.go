package main

import "fmt"

func main() {
	var n int
	total := 0

	fmt.Print("Masukkan bilangan bulat positif (n): ")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		total += i
	}

	fmt.Printf("Hasil penjumlahan dari 1 sampai %d adalah: %d\n", n, total)

}
