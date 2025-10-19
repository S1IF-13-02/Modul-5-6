package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan jumlah segitiga: ")
	fmt.Scan(&n)

	for b := 1; b <= n; b++ {
		var alas, tinggi float64
		fmt.Print("Masukkan Alas dan Tinggi Segitiga: ")
		fmt.Scan(&alas, &tinggi)

		luas := 0.5 * alas * tinggi
		fmt.Println("Luas Segitiga: ", luas)
	}
}
