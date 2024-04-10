package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u user) save() {
	fmt.Printf("save user %s", u.name)
}

func (u user) underAge() bool {
	return u.age <= 17
}

func (u *user) makeBirthDay() {
	u.age++
}

func write() {
	fmt.Println("writing")
}

func main() {
	user1 := user{"josh", 17}
	fmt.Println(user1)

	fmt.Println(user.underAge(user1))

	user1.makeBirthDay()

	fmt.Println(user.underAge(user1))
}
