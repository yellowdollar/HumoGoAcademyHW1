package main

import "fmt"

func main(){
	var a, b int

	fmt.Print("Enter the first number: ")
    fmt.Scanln(&a)
    fmt.Print("Enter the second number: ")
    fmt.Scanln(&b)

	if a > b{
		fmt.Println(a)
	}else{
		fmt.Println(b)
	}
}