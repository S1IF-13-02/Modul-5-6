package main

import "fmt"

func main() {

	var j, a, b int
	var hasil int

	fmt.Scan(&a, &b)
	hasil = 0

	for j = 1; j <= b; j += 1 {
		hasil = hasil + a
	}

	fmt.Println(hasil)

}
