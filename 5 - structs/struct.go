package main

import "fmt"

type user struct {
	name    string
	age     int
	address address
}

type address struct {
	address string
	number  int
}

type studentUser struct {
	user
	school string
}

func main() {
	fmt.Println("arquivo structs")

	var u user
	u.name = "josh"
	u.age = 21
	fmt.Println(u)

	userAdress := address{address: "teste", number: 1}

	u2 := user{"Test", 21, userAdress}
	u3 := user{name: "Test"}

	fmt.Println(u2, u3)

	u4 := user{"teste", 21, userAdress}
	ustudant := studentUser{u4, "schoolx"}

	fmt.Println(ustudant)

}
