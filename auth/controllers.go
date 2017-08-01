package auth

import (
	"encoding/json"
	"net/http"
)

// Response representation
type Response struct {
	LoggedUser LoggedUser `json:"user_info"`
	Token      string     `json:"token"`
}

const unf = "User not found"

// Auth controller
func Auth(w http.ResponseWriter, r *http.Request) {
	email, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, unf, http.StatusInternalServerError)
		return
	}
	loggedUser, err := Authenticate(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if loggedUser.ID == 0 {
		http.Error(w, unf, http.StatusNotFound)
		return
	}
	token := Token(loggedUser)
	resp := Response{
		LoggedUser: loggedUser,
		Token:      token,
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
