package main

import "fmt"

func main() {
	var a, b, hasil int
	fmt.Scan(&a, &b)
	hasil = 0

	for i := 1; i <= b; i += 1 {
		hasil = hasil + a
	}
	fmt.Println(hasil)
}
