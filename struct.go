package back

type Post struct {
	Id_post      int
	Id_user      int
	Pseudo_user  string
	Title_post   string
	Content_post string
	Comments 	 []Comment
}
type Comment struct {
	Id_post 		int
	Writer_comment 	string
	Content_comment string
}
type like struct {
	id      int
	id_post int
	effet   bool
}
type User struct {
	Id                   int
	Uusi				 string
	Age                  int
	Firstname_user       string
	Lastname_user        string
	Email_user           string
	Password_hashed_user string
	Pseudo_user          string
}

type Profil struct{ 
	Post []Post 
	User User 
}
