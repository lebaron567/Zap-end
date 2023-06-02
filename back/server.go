package main

import (
	"back"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var home = template.Must(template.ParseFiles("template/home.html"))
var registration = template.Must(template.ParseFiles("template/registration.html"))
var connexion = template.Must(template.ParseFiles("template/connexion.html"))
var explorer = template.Must(template.ParseFiles("template/Explorer.html"))
var message = template.Must(template.ParseFiles("template/message.html"))
var profil = template.Must(template.ParseFiles("template/profil.html"))
var ff = 0

type Data struct {
	Cookis string
}

func main() {
	back.InitBDD()
	http.HandleFunc("/home", Home)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/explorer", Explorer)
	http.HandleFunc("/message", Message)
	http.HandleFunc("/profil", Profil)
	http.HandleFunc("/connexion", Connexion)

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Serveur start at : http://localhost:8080/home")
	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	cookie, err2 := r.Cookie("pseudo")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/connexion", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		data = Data{Cookis: cookie.Value}
	}
	err := home.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Connexion(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var password_hashed_user string
		pseudo := r.FormValue("pseudo")
		password := r.FormValue("password")
		database := back.OpenBDD()
		err := database.QueryRow(`SELECT password_hashed_user FROM user WHERE pseudo_user = "` +pseudo+ `";`).Scan(&password_hashed_user)
		if err != nil {
			fmt.Print(err)
		}
		if back.CheckPasswordHash(password, password_hashed_user) {
			cookie := http.Cookie{
				Name:     "pseudo",
				Value:    pseudo,
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)
			fmt.Println(cookie)
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			http.Redirect(w, r, "/home", http.StatusFound)
		}

	}
	err := connexion.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Registration(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		prenom := r.FormValue("prenom")
		nom := r.FormValue("nom")
		email := r.FormValue("email")
		pseudo := r.FormValue("pseudo")
		password := r.FormValue("password")
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			log.Fatal("strconv issue")
		}
		fmt.Println(prenom, nom, email, pseudo, password, age)
		bdderr := back.AddUser(age, prenom, nom, email, password, pseudo)
		if bdderr != nil {
		} else {
			cookie := http.Cookie{
				Name:     "pseudo",
				Value:    pseudo,
				Path:     "/",
				MaxAge:   3600,
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/home", http.StatusFound)
		}
	}
	// affichier que le profile a bien été crée et rediger vers la page connexion
	err := registration.Execute(w, ff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func Explorer(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	cookie, err2 := r.Cookie("pseudo")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/connexion", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		data = Data{Cookis: cookie.Value}
	}
	err := explorer.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Message(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	cookie, err2 := r.Cookie("pseudo")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/connexion", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		data = Data{Cookis: cookie.Value}
	}
	err := message.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	cookie, err2 := r.Cookie("pseudo")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/connexion", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		data = Data{Cookis: cookie.Value}
	}
	err := profil.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
