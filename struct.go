package back

type Post struct {
	Id_post           int
	Id_user      int
	Pseudo_user  string
	Title_post   string
	Content_post string
}
type comment struct {
	id              int
	id_post         int
	id_user         int
	content_comment string
}
type like struct {
	id      int
	id_post int
	effet   bool
}
type user struct {
	id                   int
	age                  int
	firstname_user       string
	lastname_user        string
	email_user           string
	password_hashed_user string
	pseudo_user          string
}
