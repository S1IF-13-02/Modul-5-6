package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&a)

	var hasil uint64 = 1

	for i := 1; i <= a; i++ {
		hasil = hasil * uint64(i)
	}

	fmt.Println(hasil)
}