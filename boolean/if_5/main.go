package main

import "fmt"

func main() {
    var num1, num2, num3 int
    positiveCount := 0
    negativeCount := 0

    fmt.Print("Enter the first number: ")
    fmt.Scanln(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&num2)
    fmt.Print("Enter the third number: ")
    fmt.Scanln(&num3)

    if num1 > 0 {
        positiveCount++
    } else if num1 < 0 {
        negativeCount++
    }

    if num2 > 0 {
        positiveCount++
    } else if num2 < 0 {
        negativeCount++
    }

    if num3 > 0 {
        positiveCount++
    } else if num3 < 0 {
        negativeCount++
    }

    fmt.Println("Number of positive numbers:", positiveCount)
    fmt.Println("Number of negative numbers:", negativeCount)
}