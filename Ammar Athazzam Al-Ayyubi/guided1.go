package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("masukan a: ")
	fmt.Scan(&a)
	fmt.Print("masukan b: ")
	fmt.Scan(&b)

	for i := a; i <= b; i+=1 {
		fmt.Println(i, " ")
	}
	fmt.Println("selesai")
}
