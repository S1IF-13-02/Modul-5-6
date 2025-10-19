package main

import (
	"fmt"
) 

func main (){
	var n, jumlah int

	fmt.Print("Masukkan bilangan bulat = ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++{
		jumlah += i
	}
	fmt.Println(jumlah)
}