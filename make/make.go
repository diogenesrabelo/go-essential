package main

import "fmt"

func main() {

	slice1 := make([]int, 10, 11)
	fmt.Println(slice1)

	slice1 = append(slice1, 12)
	slice1 = append(slice1, 14)
	slice1 = append(slice1, 15)

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
