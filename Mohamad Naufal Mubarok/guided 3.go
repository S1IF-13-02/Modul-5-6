package main

import "fmt"

func main() {
    var a, b int
	fmt.Print("Masukan bilangan: ")
    fmt.Scan(&a, &b)

    hasil := 0
    for i := 0; i < b; i++ {
        hasil=hasil + a
    }

    fmt.Println(hasil)
}