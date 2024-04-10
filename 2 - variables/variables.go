package main

import "fmt"

func main() {
	var variable1 string = "variable1"
	variable2 := "variable 2"
	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		var3 string = "1"
		var4 string = "2"
	)

	fmt.Println(var3, var4)

	const const1 string = "const1"

	// ref memory
	var point1 int = 10
	var point2 *int = &point1

	point1++

	fmt.Println(*point2)

	point1++

	fmt.Println(*point2)
}
