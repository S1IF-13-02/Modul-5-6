package main

import "fmt"

func main() {
	var n int
	var hasil int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		hasil = hasil + i
	}
	fmt.Println("Hasil: ", hasil)
}
