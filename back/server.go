package main

import (
	"back"
	"errors"
	"fmt"
	"strings"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/google/uuid"
)

var post = template.Must(template.ParseFiles("template/post.html"))
var home = template.Must(template.ParseFiles("template/home.html"))
var registration = template.Must(template.ParseFiles("template/registration.html"))
var connexion = template.Must(template.ParseFiles("template/connexion.html"))
var explorer = template.Must(template.ParseFiles("template/Explorer.html"))
var profil = template.Must(template.ParseFiles("template/profil.html"))
var invite = template.Must(template.ParseFiles("template/invite.html"))
var ff = 0

type DataUser struct {
	Cookis string
}

type Data struct {
	User,
	Message string
	NBLike int
}

type MyData struct {
	User,
	Message string
	NBLike int
}
func main() {
	// back.AddLikeAndDislike(1, 1, 1)
	now := time.Now()
	fmt.Println("Current datetime:", now)
	back.InitBDD()
	http.HandleFunc("/post", Post)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/registration", Registration)
	http.HandleFunc("/explorer", Explorer)
	http.HandleFunc("/profil", Profil)
	http.HandleFunc("/connexion", Connexion)
	http.HandleFunc("/explorer/inviter", Inviter)
	id := uuid.New()
	fmt.Println("Generated UUID:")
	fmt.Println(id.String())

	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	fmt.Println("Serveur start at : http://localhost:8080/inviter")
	http.ListenAndServe(":8080", nil)
}

func Post(w http.ResponseWriter, r *http.Request) {
	dataUser := DataUser{}
	cookie, err2 := r.Cookie("uuid")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/inviter", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		dataUser = DataUser{Cookis: cookie.Value}
		//fmt.Println(dataUser)
	}
	if r.Method == "POST" {
		title := r.FormValue("title")
		content := r.FormValue("content")
		tag := r.FormValue("tag")
		database := back.OpenBDD()
		var id_user int
		err := database.QueryRow(`SELECT id_user FROM user WHERE uuid ="` + dataUser.Cookis + `";`).Scan(&id_user)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Println(back.AddPost(id_user, title, content, tag))
	}
	err := post.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func Home(w http.ResponseWriter, r *http.Request) {
	dataUser := DataUser{}
	posts := back.GetPosts()
	fmt.Println(posts)
	cookie, err2 := r.Cookie("uuid")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/inviter", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		dataUser = DataUser{Cookis: cookie.Value}
	}
	if r.Method == "POST"{
		like := r.FormValue("effect")
		if like == ""{
			content_comment := r.FormValue("content")
			input_id := r.FormValue("id")
			id_post, err := strconv.Atoi(input_id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			id_user := back.GetIDUserFromUUID(dataUser.Cookis)
			BDDerr := back.AddComment(id_post+1, id_user, content_comment)
			if BDDerr != nil{
				http.Error(w, BDDerr.Error(), http.StatusInternalServerError)
			}
		}
		tmp := strings.Split(like, ",")
		if tmp[0] != ""{
			fmt.Println(tmp[0])
		post_id, err := strconv.Atoi(tmp[0])
			if err != nil {
				log.Fatal(err)
			}
			user_id := back.GetIDUserFromUUID(dataUser.Cookis)
			BDDerr := back.AddLikeAndDislike(post_id, user_id, tmp[1])
			if BDDerr != nil {
				http.Error(w, BDDerr.Error(), http.StatusInternalServerError)
			}
		}
	}

	err := home.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Connexion(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var password_hashed_user string
		pseudo := r.FormValue("pseudo")
		password := r.FormValue("password")
		var uuid string
		database := back.OpenBDD()
		err := database.QueryRow(`SELECT password_hashed_user FROM user WHERE pseudo_user = "` + pseudo + `";`).Scan(&password_hashed_user)
		if err != nil {
			fmt.Print(err)
		}
		defer database.Close()
		if back.CheckPasswordHash(password, password_hashed_user) {
			err = database.QueryRow(`SELECT uuid FROM user WHERE pseudo_user = "` + pseudo + `";`).Scan(&uuid)
			if err != nil {
				fmt.Print(err)
			}
			cookie := http.Cookie{
				Name:     "uuid",
				Value:    uuid,
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
		id := uuid.New()
		//fmt.Println(prenom, nom, email, pseudo, password, age)
		bdderr := back.AddUser(id.String(), age, prenom, nom, email, password, pseudo)
		if bdderr != nil {
			fmt.Println(bdderr)
		} else {
			cookie := http.Cookie{
				Name:     "uuid",
				Value:    id.String(),
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
	err := registration.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
func Explorer(w http.ResponseWriter, r *http.Request) {
	dataUser := DataUser{}
	posts := back.GetAlPosts()
	cookie, err2 := r.Cookie("uuid")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/inviter", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		dataUser = DataUser{Cookis: cookie.Value}
	}
	if r.Method == "GET"{
		fmt.Println("dddd")
		input := r.FormValue("search")
		postsTrier := back.SearchCategorie(strings.ToLower(input))
		if postsTrier!=nil{
			posts = postsTrier
		}
		if input == ""{
			posts = back.GetAlPosts()
		}
	}
	if r.Method == "POST" {
		input := r.FormValue("effect")
		tmp := strings.Split(input, ",")
		post_id , err := strconv.Atoi(tmp[0])
		if err != nil {
			log.Fatal(err)
		}
		user_id := back.GetIDUserFromUUID(dataUser.Cookis)
		BDDerr := back.AddLikeAndDislike(post_id, user_id, tmp[1])
		if BDDerr != nil {
			http.Error(w, BDDerr.Error(), http.StatusInternalServerError)
		}

	}
	err := explorer.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Profil(w http.ResponseWriter, r *http.Request) {
	var user back.User
	var userModif back.User
	var data back.Profil
	uuid, err2 := r.Cookie("uuid")
	post := back.GetAlPostsUser(uuid.Value)
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):
			http.Redirect(w, r, "/inviter", http.StatusFound)
		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		user = back.GetUser(uuid.Value)
	}
	if r.Method == "POST" {
		userModif = user
		if strings.Contains(r.FormValue("Prenon"), " ") {
			userModif.Firstname_user = user.Firstname_user
		} else {
			userModif.Firstname_user = r.FormValue("Prenon")
		}
		if strings.Contains(r.FormValue("nom"), " ") {
			userModif.Lastname_user = user.Lastname_user
		} else {
			userModif.Lastname_user = r.FormValue("nom")
		}
		if strings.Contains(r.FormValue("pseudo"), " ") {
			userModif.Pseudo_user = user.Pseudo_user
		} else {
			userModif.Pseudo_user = r.FormValue("pseudo")
		}
		userModif.Age, _ = strconv.Atoi(r.FormValue("age"))
		if userModif.Firstname_user != "" && userModif.Lastname_user != "" && userModif.Pseudo_user != "" {
			fmt.Println(back.UpdateUser(userModif))
		}
	}
	data.Post = post
	data.User = user
	err := profil.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Inviter(w http.ResponseWriter, r *http.Request) {
	dataUser := DataUser{}
	cookie, err2 := r.Cookie("uuid")
	if err2 != nil {
		switch {
		case errors.Is(err2, http.ErrNoCookie):

		default:
			log.Println(err2)
			http.Error(w, "server error", http.StatusInternalServerError)
		}
		return
	} else {
		dataUser = DataUser{Cookis: cookie.Value}
		//http.Redirect(w, r, "/home", http.StatusFound)
	}
	err := invite.Execute(w, dataUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
