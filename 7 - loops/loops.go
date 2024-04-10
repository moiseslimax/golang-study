package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("pass 1 sec", i)
		time.Sleep(time.Second)
	}

	for j := 0; j < 10; j++ {
		fmt.Println("pass 1 sec", j)
		time.Sleep(time.Second)
	}

	names := []string{"josh", "maria", "joao"}

	for _, name := range names {
		fmt.Println(name)
	}

	for _, word := range "names" {
		fmt.Println(string(word))
	}
}
