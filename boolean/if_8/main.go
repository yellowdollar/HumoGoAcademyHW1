package main

import "fmt"

func main() {
    var num1, num2 int

    fmt.Print("Input first number: ")
    fmt.Scanln(&num1)
	
    fmt.Print("Input second number: ")
    fmt.Scanln(&num2)

    if num1 > num2 {
        fmt.Println("Bigger:", num1)
        fmt.Println("Less:", num2)
    } else if num2 > num1 {
        fmt.Println("Bigger:", num2)
        fmt.Println("Less:", num1)
    } else {
        fmt.Println("Equal")
    }
}