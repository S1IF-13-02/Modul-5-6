package main

import (
	"fmt"
)

func main() {
	for{
	var n int
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Print("selesai")
		break
	}
	sum := 0
	for i := 1; i <= n; i++ {	
		sum = sum + i
	}
	fmt.Println("hasil:", sum)
   }
}
