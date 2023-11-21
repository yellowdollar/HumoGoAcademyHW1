package main

import "fmt"

func main() {
    var A, B int

    fmt.Print("Введите A: ")
    fmt.Scanln(&A)

    fmt.Print("Введите zB: ")
    fmt.Scanln(&B)

    if A != B {
        A = A + B
        B = A
    } else {
        A = 0
        B = 0
    }

    fmt.Println("Новое значение переменной A:", A)
    fmt.Println("Новое значение переменной B:", B)
}
