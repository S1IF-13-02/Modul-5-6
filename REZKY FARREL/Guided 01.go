package main

import "fmt"

	func main() {
 	var a, b int
 	var j int
 	fmt.Print("Masukan bilangan a: ")
 	fmt.Scan(&a)
 	fmt.Print("Masukan bilangan b: ")
 	fmt.Scan(&b)
 	for j = a; j <=b; j+=1 {
 	fmt.Print(j," ")
 }
}