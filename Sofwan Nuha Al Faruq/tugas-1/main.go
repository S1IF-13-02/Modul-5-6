package main

import "fmt"

func main() {
    var a, hasil int

    fmt.Print("Masukkan bilangan a: ")
    fmt.Scan(&a)
    hasil = 0
    for i := 1; i <= a; i++ {
        hasil += i
    }
    fmt.Println("Hasil penjumlahan dari 1 sampai", a, "adalah:", hasil)
}