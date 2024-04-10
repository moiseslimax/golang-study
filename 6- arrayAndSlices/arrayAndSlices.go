package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("array and splices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]string{"1", "2", "3", "4", "5"}
	fmt.Println(array3)

	slice := []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, "18")
	fmt.Println(slice)

	//internal arrays
	slice2 := make([]int, 10, 15)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
}
