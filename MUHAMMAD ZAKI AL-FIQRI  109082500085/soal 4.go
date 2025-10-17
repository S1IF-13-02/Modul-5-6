package main

import (
	"fmt"
)

func main() {
	var n int
	var hasil int = 1

	fmt.Print("masukan faktorial : ")
	fmt.Scan(&n)

	for j := 1; j <= n; j++ {

		hasil = hasil * j

	}
	fmt.Print(hasil)
}
