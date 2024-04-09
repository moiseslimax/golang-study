package main

import (
	"fmt"
	"module/utils"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	utils.Write()
	utils.Write2()

	erro := checkmail.ValidateFormat("moiseslimax@gmail.com")
	fmt.Println(erro)
}
