package db

import (
	"testing"

	"github.com/Jonny-Burkholder/gorilla-mux-example-api/internal/user"
)

var testUser = user.NewUser("001", "Jeff", "Bezos")

var testDB = &DB{
	DataBase: []*user.User{testUser},
}

func TestGetUserByID(t *testing.T) {
	users := testDB.GetUserByID("001")
	if len(users) != 1 {
		t.Error("Invalid result")
	}
	if users[0] != testUser {
		t.Errorf("Wanted %v, got %v", testUser, users[0])
	}
}

func TestGetUserByFirstName(t *testing.T) {
	users := testDB.GetUserByFirstName("Jeff")
	if len(users) != 1 {
		t.Error("Invalid result")
	}
	if users[0] != testUser {
		t.Errorf("Wanted %v, got %v", testUser, users[0])
	}
}

func TestGetUserByLastName(t *testing.T) {
	users := testDB.GetUserByLastName("Bezos")
	if len(users) != 1 {
		t.Error("Invalid result")
	}
	if users[0] != testUser {
		t.Errorf("Wanted %v, got %v", testUser, users[0])
	}
}

func TestPutUser(t *testing.T) {
	user := user.NewUser("002", "Bruce", "Banner")
	testDB.PutUser(user)

	if testDB.DataBase[1] != user {
		t.Errorf("Wanted %v, got %v", testUser, testDB.DataBase[1])
	}
}
