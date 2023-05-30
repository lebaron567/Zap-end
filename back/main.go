package main

import (
	"back"
	"fmt"
	"html/template"
	"net/http"
)

var home = template.Must(template.ParseFiles("template/home.html"))
var registration = template.Must(template.ParseFiles("template/registration.html"))
var explorer = template.Must(template.ParseFiles("template/Explorer.html"))
var message = template.Must(template.ParseFiles("template/message.html"))
var profil = template.Must(template.ParseFiles("template/profil.html"))
var ff = 0

func main() {
	back.InitBDD()
	// value := back.GetAllFeildInTable("users")
	// fmt.Println(value)

	// back.AddUser(18, "b", "b", "a@y.c", "secret", "a")
	value := back.GetAllFeildInTable("users")
	fmt.Println(value)
	fmt.Println(value[0][6])
	fmt.Println(value[1][6])

	http.HandleFunc("/home", back.Home)
	http.HandleFunc("/registration", back.Registration)
	http.HandleFunc("/explorer", back.Explorer)
	http.HandleFunc("/message", back.Message)
	http.HandleFunc("/profil", back.Profil)
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Serveur start at : http://localhost:8080/home")
	http.ListenAndServe(":8080", nil)
}
