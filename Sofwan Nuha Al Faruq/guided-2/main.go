package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah segitiga: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		var alas, tinggi float64
		fmt.Printf("Masukkan alas dan tinggi segitiga ke-%d: ", i)
		fmt.Scan(&alas, &tinggi)

		luas := 0.5 * alas * tinggi
		fmt.Printf("Luas segitiga ke-%d = %.2f\n", i, luas)
	}
}
