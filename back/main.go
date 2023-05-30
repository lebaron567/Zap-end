package main

import (
	"back"
	"fmt"
)

func main() {
	back.InitBDD()
	value := back.GetAllFeildInTable("users")
	fmt.Println(value)

	// back.AddUser(18, "arthur", "a", "a@y.c", "secret2", "a")
	// value = back.GetAllFeildInTable("users")
	// fmt.Println(back.CheckPasswordHash("secret", value[5]))
	// fmt.Println(value)
}
