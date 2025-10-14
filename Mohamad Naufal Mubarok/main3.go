package main
import (
		"fmt"
	)

func main() {
	var a, b int
	fmt.Print("Masukan bilangan a: ")
	fmt.Scan(&a)

	fmt.Print("masukan bilangan b: ")
	fmt.Scan(&b)

	hasil := 1
	for i := 1; i <=b; i++ {
		hasil *= a
	} 
		fmt.Println(hasil)
}