package main

import "fmt"

func main() {
	var alas float64
	var tinggi float64
	var n int
	fmt.Print("Masukkan banyaknya segitiga: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("Segitiga ke-%d", i)
		fmt.Print("\nMasukkan nilai alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan nilai tinggi: ")
		fmt.Scan(&tinggi)
		luas := 0.5 * alas * tinggi
		fmt.Println("Luas segitiga: ", luas)
	}
	fmt.Println("Perhitungan selesai")
}
