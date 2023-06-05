package back

type post struct {
	id           int
	id_user      int
	title_post   string
	content_post string
	pseudo_user  string
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
