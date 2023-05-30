package main

import (
	"back"
	"fmt"
)

func main() {
	back.InitBDD()
	// value := back.GetAllFeildInTable("users")
	// fmt.Println(value)

	// back.AddUser(18, "b", "b", "a@y.c", "secret", "a")
	value := back.GetAllFeildInTable("users")
	fmt.Println(value)
	fmt.Println(value[0][6])
	fmt.Println(value[1][6])

}
