package main

import "fmt"

func recoveryExec() {
	if r := recover(); r != nil {
		fmt.Println("recovery exec")
	}

}

func aproveStudent(n1, n2 float64) bool {
	defer recoveryExec()
	med := (n1 + n2) / 2

	if med > 6 {
		return true
	} else if med < 6 {
		return false
	}

	panic("med = 6")
}

func main() {
	fmt.Println(aproveStudent(6, 6))
	fmt.Println("end exec")

}
