package main

import "fmt"

func main() {

	var small_number int
	var big_number int
	fmt.Println("Enter a Big number")
	fmt.Scan(&big_number)
	fmt.Println("Enter a small number")
	fmt.Scan(&small_number)
	fmt.Println(big_number % small_number)

}
