package main

import "fmt"

func main() {
	var bilangan, i, total int
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&bilangan)

	for i = 1; i <= bilangan; i++ {
		total = total + i
	}
	fmt.Println("Total: ", total)
}
