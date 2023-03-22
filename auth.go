package authenticate

import (
	"fmt"
	"net/http"
)

type user struct {
	Username string
	Password string
}

var users = []user{
	{Username: "alice", Password: "password1"},
	{Username: "bob", Password: "password2"},
}

func authenticate(username string, password string) bool {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	if !authenticate(username, password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}

	// add connection to list of active connections
	connections[conn] = true
}
