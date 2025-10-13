package main

import "fmt"
func main() {
	var x, y int
	var hasil int = 1
	fmt.Print("Masukan angka pokok   : ")
	fmt.Scan(&x)
	fmt.Print("Masukan angka pangkat : ")
	fmt.Scan(&y)
	for i := 1; i <= y; i++ {
		hasil = hasil * x
	}
	fmt.Println("Hasil dari", x, "pangkat", y, "adalah", hasil)
}