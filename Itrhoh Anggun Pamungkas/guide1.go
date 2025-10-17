package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("masukkan a: ")
	fmt.Scan(&a)
	fmt.Print("masukkan b: ")
	fmt.Scan(&b)

	for i := a ; i <= b ; i++{
		fmt.Print(i, " ")
	}
}