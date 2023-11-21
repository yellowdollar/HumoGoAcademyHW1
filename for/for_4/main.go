package main

import "fmt"

func main() {
	var price float64

	fmt.Println("Input price of 1 kg: ")
	fmt.Scan(&price)

	for i := 1; i <= 10; i++ {
		fmt.Println(price * float64(i))
	}
}
