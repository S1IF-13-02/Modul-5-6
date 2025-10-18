package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a, &b)
	var hasil int = 0

	for i := 0; i < b; i++ {
		hasil += a
	}
	fmt.Println(hasil)
}
