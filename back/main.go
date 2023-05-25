package main

import (
	"back"
	"fmt"
)

func main() {
	
	value := back.Get("*", "users")
	fmt.Println(value)

	back.AddUser(18, "a", "a", "a@y.c", "secret2", "b")
	value = back.Get("*", "users")
	// fmt.Println(back.CheckPasswordHash("secret", value[5]))
	fmt.Println(value)
	back.Serveur()
}