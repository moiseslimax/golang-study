package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func mathCalcs(n1, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

// func named return
func mathCalcsNamed(n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func printEnd() {
	fmt.Println("finish main function")
}

func main() {
	defer printEnd()

	sumx := sum(10, 30)
	fmt.Println(sumx)

	var f = func() {
		fmt.Println("func f ")
	}

	f()

	resSum, _ := mathCalcs(10, 20)

	fmt.Println(resSum)

	mathCalcsNamed(1, 2)
}
