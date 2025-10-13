package main

import "fmt"

func main() {
	var n int
	var init int

	fmt.Printf("Masukkan perulangan: ")
	fmt.Scanln(&n)
	fmt.Printf("Masukkan awal: ")
	fmt.Scanln(&init)

	for i := init ; i <= n ; i++ {	
			fmt.Print(i, " ")
	}
}