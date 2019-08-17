package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route all routes of web
func Route() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", Index)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	return router

}
