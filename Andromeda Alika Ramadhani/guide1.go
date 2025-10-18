package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai awal: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai akhir: ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}
}
