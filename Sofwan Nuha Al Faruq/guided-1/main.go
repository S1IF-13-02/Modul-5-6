package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan Nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan Nilai b : ")
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
}
