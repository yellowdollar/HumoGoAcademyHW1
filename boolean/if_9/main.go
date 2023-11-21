package main

import "fmt"

func main() {
    var A, B, C float64

    fmt.Print("Введите значение переменной A: ")
    fmt.Scanln(&A)
    fmt.Print("Введите значение переменной B: ")
    fmt.Scanln(&B)

    if A > B {
        C = A
		A = B
		B = C
    }

    fmt.Println("Новое значение переменной A:", A)
    fmt.Println("Новое значение переменной B:", B)
}