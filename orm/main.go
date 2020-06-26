package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	fmt.Println("Go ORM Tutorial")

	handleRequests()

	InitialMigration()
}
