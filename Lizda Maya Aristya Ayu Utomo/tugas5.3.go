package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("masukan dua bilangan: ")
	fmt.Scan(&a, &b)

	hasil := 1
	for i := 1; i <= b; i++ {
		hasil *= a

	}

	fmt.Println(hasil)

}
