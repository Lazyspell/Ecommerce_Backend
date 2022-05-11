package models

type Users struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type GoogleObject struct {
	ID        string
	Email     string
	Verified  bool
	Name      string
	GivenName string
	Picture   string
	Locale    string
}

//TODO: need to remove the Password frield from teh User model and switch everything to the password model
type Password struct {
	Password string `json:"password"`
}
