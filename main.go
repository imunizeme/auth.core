package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/imunizeme/auth.core/auth"
	config "github.com/imunizeme/config.core"
	"github.com/nuveo/log"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}
	n := negroni.Classic()
	router := mux.NewRouter()
	n.Use(cors.Default())
	auth.RouterRegister(router)
	n.UseHandler(router)
	n.Run(fmt.Sprintf("%s:%d", config.Get.Auth.Host, config.Get.Auth.Port))
}
