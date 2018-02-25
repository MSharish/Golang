package main

import "fmt"

func main() {

	var read_name string
	fmt.Println("Enter you name")
	fmt.Scan(&read_name)
	fmt.Println("Hi " + read_name + " Welcome to Go lang programming")
}
