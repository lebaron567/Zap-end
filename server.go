package main

import (
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", Home)
	go exec.Command("cmd", "/C", "start", "http://localhost:8080").Run()
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/home.html")
}
