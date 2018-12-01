package main

import "github.com/gorilla/mux"

func registerRoutes(router *mux.Router) {
	router.HandleFunc("/jobs", jobsGetHandler).Methods("GET")
	// router.HandleFunc("/jobs", jobsPostHandler).Methods("POST")
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	registerRoutes(router)
	return router
}
