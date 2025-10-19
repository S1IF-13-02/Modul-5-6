package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&n)

	total := 0
	for i := 1; i <= n; i++ {
		total = total + i

	}
	fmt.Println("Hasil penjumlahan dari 1 sampai", n, "adalah", total)

}
