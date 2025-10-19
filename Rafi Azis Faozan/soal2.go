package main

import "fmt"

func main() {
	var n int
	var jari2 float64
	var tinggi float64
	fmt.Print("Masukkan banyaknya kerucut: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println("Volume kerucut ke-", i)
		fmt.Print("Masukkan nilai jari2: ")
		fmt.Scan(&jari2)
		fmt.Print("Masukkan nilai tinggi: ")
		fmt.Scan(&tinggi)
		pi := 3.14
		Volume := (1.0 / 3.0) * pi * jari2 * jari2 * tinggi
		fmt.Println("Volume kerucutt: ", Volume)
	}
}
