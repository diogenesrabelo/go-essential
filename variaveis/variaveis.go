package main

import "fmt"

func main() {
	var name string = "Diogenes"
	lastName := "Rabelo"

	var (
		years     int  = 27
		isMarried bool = true
	)

	const gender string = "Masculino"
	fmt.Println(name)
	fmt.Println(lastName)
	fmt.Println(years)
	fmt.Println(isMarried)
	fmt.Println(gender)
}
