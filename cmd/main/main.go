package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Jonny-Burkholder/gorilla-mux-example-api/internal/db"
	"github.com/Jonny-Burkholder/gorilla-mux-example-api/internal/user"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	database.PutUser(user.NewUser("001", "jeff", "bezos"))
	database.PutUser(user.NewUser("002", "bruce", "banner"))
	database.PutUser((user.NewUser("003", "emma", "watson")))

	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUserID).Methods("GET")
	router.HandleFunc("/users", putUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))

}

var database = db.NewDB()

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//we'll assume one query at a time, for simplicity
	log.Println("Endpoint reached: get users")
	vars := r.URL.Query()
	if len(vars) > 1 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	} else if len(vars) < 1 {
		json.NewEncoder(w).Encode(database.DataBase)
		return
	} else {
		var res []*user.User
		if id := vars.Get("id"); id != "" {
			res = database.GetUserByID(id)
		} else if first := vars.Get("first"); first != "" {
			res = database.GetUserByFirstName(first)
		} else if last := vars.Get("last"); last != "" {
			res = database.GetUserByLastName(last)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(res)
	}
}

func getUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Endpoint reached - get user by id")
	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok != false {
		res := database.GetUserByID(id)
		json.NewEncoder(w).Encode(res)
		return
	}
	http.Error(w, "Invalid request", http.StatusBadRequest)
}

func putUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Endpoint reached - Put User")

	user := &user.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
	}

	database.PutUser(user)

	log.Println("User successfully updated")
}
