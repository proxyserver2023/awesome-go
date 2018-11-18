package simpleapi

import "github.com/gorilla/mux"

func RegisterRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", Home)
	registerAuthRoutes(router)

	return router
}

func registerAuthRoutes(router *mux.Router) {
	router.HandleFunc("/auth", Login)
}
