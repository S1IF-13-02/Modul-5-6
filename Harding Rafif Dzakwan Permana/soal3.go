package main

import "fmt"

func main () {
	var a  int 
	var hasil int

	fmt.Scan(&a)

	hasil = 1 

	for i := 1; i <= a; i++ {
		hasil = hasil * a
	}
	fmt.Println (hasil)
}