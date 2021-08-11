package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Array & Slices")

	var array1 [5]uint8

	array1[0] = 1
	array1[1] = 5
	array1[2] = 3
	array1[3] = 4
	array1[4] = 2

	fmt.Println(array1)

	slice := []int{}

	slice = append(slice, 1)
	slice = append(slice, 5)
	slice = append(slice, 3)
	slice = append(slice, 2)
	slice = append(slice, 4)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Print(slice)
}
