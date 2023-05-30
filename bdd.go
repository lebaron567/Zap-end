package back

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

func OpenBDD() *sql.DB {
	database, bdderr := sql.Open("sqlite3", "./BDD.db")
	if bdderr != nil {

		log.Fatal(bdderr.Error())

	}
	return database
}

func InitBDD() {
	database := OpenBDD()
	defer database.Close()
	tmp := `
	CREATE TABLE IF NOT EXISTS "user" (
		"id"					INTEGER NOT NULL UNIQUE,
		"age"					INTEGER NOT NULL,
		"firstname_user"		VARCHAR(20) NOT NULL,
		"lastname_user"			VARCHAR(30) NOT NULL,
		"email_user"			VARCHAR(50) NOT NULL UNIQUE,
		"password_hashed_user"	VARCHAR(45) NOT NULL,
		"pseudo_user"	V		ARCHAR(20) NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
		
	);
	`

	_, bdderr := database.Exec(tmp)
	if bdderr != nil {
		log.Fatal(bdderr.Error())
	}

}

func GetAllFeildInTable(table string) [][]string {

	var result [][]string
	database := OpenBDD()
	rows, err := database.Query("SELECT * FROM " + table)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var feildOfTable []string
		var id int
		var age int
		var firstname string
		var lastname string
		var email string
		var password string
		var pseudo string
		err = rows.Scan(&id, &age, &firstname, &lastname, &email, &password, &pseudo)
		if err != nil {
			log.Fatal(err)
		}
		feildOfTable = append(feildOfTable, strconv.Itoa(id), strconv.Itoa(age), firstname, lastname, email, password, pseudo)
		result = append(result, feildOfTable)

	}
	return result
}

func GetFeildInTable(feild string, table string) []string {
	var feildOfTable []string
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
			feildOfTable = append(feildOfTable, strconv.Itoa(id), strconv.Itoa(age), firstname, lastname, email, password, pseudo)
		}
		if feild == "id" {
			err = rows.Scan(&id)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, strconv.Itoa(id))

		}
		if feild == "age" {
			err = rows.Scan(&age)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, strconv.Itoa(age))

		}
		if feild == "firstname" {
			err = rows.Scan(&firstname)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, firstname)
		}
		if feild == "lastname" {
			err = rows.Scan(&lastname)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, lastname)
		}
		if feild == "email" {
			err = rows.Scan(&email)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, email)
		}
		if feild == "password" {
			err = rows.Scan(&password)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, password)
		}
		if feild == "pseudo" {
			err = rows.Scan(&pseudo)
			if err != nil {
				log.Fatal(err)
			}
			feildOfTable = append(feildOfTable, pseudo)
		}

	}
	return feildOfTable
}

func AddUser(age int, firstname string, lastname string, email string, password string, pseudo string) {
	database := OpenBDD()
	password = HashPassword(password)
	fmt.Println(password)
	if age < 13 {
		log.Fatal("grandis et tu pourras parler")
	}
	statement, BDDerr := database.Prepare(`INSERT INTO user(age, firstname_user, lastname_user, email_user, password_user, pseudo_user) VALUES(?,?,?,?,?,?);`)
	if BDDerr != nil {
		log.Fatal("ERROR 500 : error Prepare new user \n ", BDDerr)
	}
	_, BDDerr = statement.Exec(strconv.Itoa(age), firstname, lastname, email, password, pseudo)
	if BDDerr != nil {
		log.Fatal("ERROR 500 : error exec new user \n ", BDDerr)
	}
	defer database.Close()
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
