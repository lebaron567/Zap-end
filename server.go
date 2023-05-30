package back

import (
	"fmt"
	"net/http"
	"text/template"
)

var home = template.Must(template.ParseFiles("template/home.html"))
var registration = template.Must(template.ParseFiles("template/registration.html"))
var explorer = template.Must(template.ParseFiles("template/Explorer.html"))
var message = template.Must(template.ParseFiles("template/message.html"))
var profil = template.Must(template.ParseFiles("template/profil.html"))
var ff = 0

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/explorer", Explorer)
	http.HandleFunc("/message", Message)
	http.HandleFunc("/profil", Profil)
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Serveur start at : http://localhost:8080/home")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := home.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Registration(w http.ResponseWriter, r *http.Request)  {
	prenom := r.FormValue("prenom")
	nom := r.FormValue("nom")
	email := r.FormValue("email")
	pseudo := r.FormValue("pseudo")
	password := r.FormValue("password")
	age := r.FormValue("age")
	fmt.Println(prenom,nom,email,pseudo,password,age)
	err := registration.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Explorer(w http.ResponseWriter, r *http.Request) {
	err := explorer.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Message(w http.ResponseWriter, r *http.Request) {
	err := message.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	err := profil.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}