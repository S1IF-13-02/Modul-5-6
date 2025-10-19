package main

import "fmt"

func main() {
    var n int
	fmt.Print("Input n: ")
	fmt.Scan(&n)
	total := 0
	for i := 1; i <= n; i++ {
		total = total + i
	}
	fmt.Print("Output: ", total)
}