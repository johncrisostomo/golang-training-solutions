package main

import "fmt"

func main() {
	var smaller, larger int

	fmt.Print("Enter a small number: ")
	fmt.Scan(&smaller)
	fmt.Print("Enter a larger number: ")
	fmt.Scan(&larger)

	fmt.Printf("%d %% %d = %d \n", larger, smaller, (larger % smaller))
}
