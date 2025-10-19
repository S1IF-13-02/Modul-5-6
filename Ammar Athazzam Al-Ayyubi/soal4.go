package main

import (
	"fmt"
)

func main() {
	for {
		var n int
		fmt.Print("Bilangan Non(-): ")
		fmt.Scan(&n)

		if n < -1 {
			fmt.Printf("selesai.")
			break
		}
		factorial := 1
		for i := 1; i <= n; i++ {
			factorial = factorial * i
		}
		fmt.Printf("Hasil Factorial: (%d)\n\n", factorial)
	}
}
