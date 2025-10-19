package main

import (
	"fmt"
)

func main() {
	var n, hasil int
	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&n)

	hasil = 1
	for j := 1; j <= n; j++{
		hasil *=j
	} 
	fmt.Println(hasil)

}