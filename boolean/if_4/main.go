package main

import "fmt"

func main(){
	var a, b, c int

	fmt.Println("Input first number: ")
	fmt.Scan(&a)

	fmt.Println("Input second number: ")
	fmt.Scan(&b)

	fmt.Println("Input third number: ")
	fmt.Scan(&c)

	if a > 0 && b > 0 && c > 0 {
		fmt.Println(3)
	}else if (a > 0 && b > 0 && c < 0) || (a > 0 && b < 0 && c > 0) || (a < 0 && b > 0 && c > 0){
		fmt.Println(2)
	}else if (a > 0 && b < 0 && c < 0) || (a < 0 && b > 0 && c < 0) || (a < 0 && b < 0 && c > 0){
		fmt.Println(1)
	}else if a < 0 && b < 0 && c < 0{
		fmt.Println(0)
	}
}