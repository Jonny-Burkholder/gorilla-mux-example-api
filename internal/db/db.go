package db

import (
	"github.com/Jonny-Burkholder/gorilla-mux-example-api/internal/user"
)

//DB stores users
type DB struct {
	DataBase []*user.User `json:"database"`
}

//NewDB returns a pointer to a new database
//object with an instantiated, but empty
//db.DataBase field
func NewDB() *DB {
	return &DB{
		DataBase: make([]*user.User, 0),
	}
}

//GetUserByFirstName takes a string argument
//and returns a slice of all users matching
//that name. If no users are found, it returns
//an empty slice
func (db *DB) GetUserByFirstName(name string) []*user.User {
	res := make([]*user.User, 0)
	for _, user := range db.DataBase {
		if user.First == name {
			res = append(res, user)
		}
	}
	return res
}

//GetUserByLastname takes a string argument
//and returns a slice of all users matching
//that name. If no users are found, it returns
//an empty slice
func (db *DB) GetUserByLastName(name string) []*user.User {
	res := make([]*user.User, 0)
	for _, user := range db.DataBase {
		if user.Last == name {
			res = append(res, user)
		}
	}
	return res
}

//GetUserByLastname takes a string argument
//and returns a slice of all users matching
//that name. If no users are found, it returns
//an empty slice
func (db *DB) GetUserByID(id string) []*user.User {
	res := make([]*user.User, 0)
	for _, user := range db.DataBase {
		if user.ID == id {
			res = append(res, user)
		}
	}
	return res
}

//PutUser takes a pointer to user.User. If
//the user already exists in the database,
//the user information is updated. Otherwise,
//the new user is appended to the database
func (db *DB) PutUser(user *user.User) {
	for i, u := range db.DataBase {
		if u.ID == user.ID {
			db.DataBase[i] = user
		}
	}
	db.DataBase = append(db.DataBase, user)
}
