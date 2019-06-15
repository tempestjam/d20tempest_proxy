package main

import (
	"d20tempest_authentificator/controllers"
	fmt "fmt"
	http "net/http"

	mux "github.com/gorilla/mux"
)

func createRouter() http.Handler {
	router := mux.NewRouter()

	//fmt.Println("load GET /hello route")
	//router.HandleFunc("/hello", SendHello).Methods("GET")

	//fmt.Println("load POST /session/login route")
	//router.HandleFunc("/session/login", controllers.LogIn).Methods("POST")

	fmt.Println("load WS /party/join route")
	router.HandleFunc("/p", controllers.JoinParty).Methods("GET")

	return router
}
