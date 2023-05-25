package back

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

type users struct {
	id        int
	firstname string
	lastname  string
	email     string
	password  string
	pseudo    string
}

func OpenBDD() *sql.DB {
	database, bdderr := sql.Open("sqlite3", "./BDD.db")
	if bdderr != nil {

		log.Fatal(bdderr.Error())

	}
	return database
}

func InitBDD() {
	// var test users
	database := OpenBDD()
	defer database.Close()
	tmp := `
	CREATE TABLE IF NOT EXISTS "users" (
		"id"	INTEGER NOT NULL UNIQUE,
		"age"	INTEGER NOT NULL UNIQUE,
		"firstname"	TEXT NOT NULL,
		"lastname"	TEXT NOT NULL,
		"email"		CHAR(50) NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		"pseudo"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
		
	);
	`

	_, bdderr := database.Exec(tmp)
	if bdderr != nil {
		log.Fatal(bdderr.Error())
	}

}

func Get(feild string, table string) []string {
	var users []string
	database := OpenBDD()
	rows, err := database.Query("SELECT " + feild + " FROM " + table)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id int
		var age int
		var firstname string
		var lastname string
		var email string
		var password string
		var pseudo string
		if feild == "*" {
			err = rows.Scan(&id, &age, &firstname, &lastname, &email, &password, &pseudo)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, strconv.Itoa(id), strconv.Itoa(age), firstname, lastname, email, password, pseudo)
		}
		if feild == "id" {
			err = rows.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, strconv.Itoa(id))

		}
		if feild == "age" {
			err = rows.Scan(&age)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, strconv.Itoa(age))

		}
		if feild == "firstname" {
			err = rows.Scan(&firstname)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, firstname)
		}
		if feild == "lastname" {
			err = rows.Scan(&lastname)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, lastname)
		}
		if feild == "email" {
			err = rows.Scan(&email)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, email)
		}
		if feild == "password" {
			err = rows.Scan(&password)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, password)
		}
		if feild == "pseudo" {
			err = rows.Scan(&pseudo)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, pseudo)
		}

	}
	return users
}

func AddUser(age int, firstname string, lastname string, email string, password string, pseudo string) {
	database := OpenBDD()
	password = HashPassword(password)
	fmt.Println(password)
	if age < 13 {
		log.Fatal("grandis et tu pourras parler")
	}
	statement, BDDerr := database.Prepare(`INSERT or IGNORE INTO users(age, firstname, lastname, email, password, pseudo) VALUES(?,?,?,?,?,?);`)
	if BDDerr != nil {
		log.Fatal("ERROR 500 : error Prepare new user \n ", BDDerr)
	}
	statement.Exec(strconv.Itoa(age), firstname, lastname, email, password, pseudo)
}

func HashPassword(password string) string {
	bytes, hashingErr := bcrypt.GenerateFromPassword([]byte(password), 14)
	if hashingErr != nil {
		log.Fatal(hashingErr)
	}
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	hashingErr := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return hashingErr == nil
}
