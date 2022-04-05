package user

//user holds user information
type User struct {
	ID    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
}

//NewUser takes 3 string arguments and returns
//a pointer to a new user
func NewUser(id, first, last string) *User {
	return &User{id, first, last}
}
