package main

import "fmt"


const PI = 3.141592653589793 

func main() {
	var n int
	var r, t float64 
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&r, &t) 
		volume := (PI * r * r * t) / 3.0
		fmt.Println(volume) 
	}
}
