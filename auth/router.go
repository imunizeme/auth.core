package auth

import (
	"github.com/gorilla/mux"
)

func routers(r *mux.Router) {
	r.HandleFunc("/auth", Auth).Methods("POST")
}

// RouterRegister auth app
func RouterRegister(r *mux.Router) {
	routers(r)
}
