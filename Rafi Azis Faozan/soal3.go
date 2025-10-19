package main

import "fmt"

func main() {
	var a int
	var b int
	var hasil int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	hasil = 1
	for i := 1; i <= b; i++ {
		hasil = hasil * a
	}
	fmt.Println("Hasil: ", hasil)
}
