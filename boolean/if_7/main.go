package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Введите первое число: ")
    fmt.Scanln(&num1)
    fmt.Print("Введите второе число: ")
    fmt.Scanln(&num2)

    if num1 < num2 {
        fmt.Println("Порядковый номер меньшего числа: 1")
    } else if num2 < num1 {
        fmt.Println("Порядковый номер меньшего числа: 2")
    } else {
        fmt.Println("Числа равны")
    }
}