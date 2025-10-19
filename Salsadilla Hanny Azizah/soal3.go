package main

import (
	"fmt"
)

func main(){
	var n, x int
	fmt.Print("Masukkan bilangan = ")
	fmt.Scan(&n, &x)

	hasil := 1
	for j := 1; j <= x; j++{
		hasil *= n
	}
	fmt.Println(hasil)
}