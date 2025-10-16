package main

import "fmt"

func main() {
	var n, i, total int
	fmt.Scan(&n)

	total = 0
	for i = 1; i <= n; i++ {
		total = total + i
	}

	fmt.Println(total)
}
