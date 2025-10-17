package main

import (
	"fmt"
)

func main() {
	var a, n int
	var hasil int = 1

	fmt.Print("masukan nilai pangkat : ")
	fmt.Scan(&n)
	fmt.Print("masukan nilai basis : ")
	fmt.Scan(&a)

	for j := 1; j <= n; j++ {

		hasil = hasil * a

	}
	fmt.Print(hasil)

}
