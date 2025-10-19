package main

import "fmt"

func main() {
	var n int
	var hasil int

	fmt.Print(" masukan bilagan n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		hasil += i
	}

	fmt.Println(" keluaran : ", hasil)

}
