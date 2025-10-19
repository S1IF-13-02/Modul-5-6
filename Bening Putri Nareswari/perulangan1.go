package main

import "fmt"

func main() {
	var a, b int

	fmt.Println("Masukkan dua bilangan bulat a dan b: ")
	fmt.Scan(&a, &b)

	for i := a; i <= b; i++ {
		fmt.Print(i, "")
	}
	fmt.Println()
}
