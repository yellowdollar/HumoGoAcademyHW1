package main

import "fmt"

func main(){
	var a int

	fmt.Println("Input first number: ")
	fmt.Scan(&a)

	if(a > 0){
		fmt.Println(a + 1)
	}else{
		fmt.Println(a - 2)
	}
}