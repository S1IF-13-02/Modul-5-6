package main

import "fmt"

func main() {
	var n int
	var hj int

	fmt.Print("masukan n : ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hj = hj + i
	}
	fmt.Println("hasil : ", hj)
}
