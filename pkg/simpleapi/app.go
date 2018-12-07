package simpleapi

import (
	"fmt"
	"net/http"
	"os"
	"path"

	"github.com/joho/godotenv"
)

func Run() {
	// REGISTER ROUTER
	router := RegisterRouter()

	// GET DB
	db := GetDB()
	perr := db.Ping()

	if perr != nil {
		fmt.Println("Could Not Connect To Database")
		fmt.Println(perr)
	}

	// Load all ENV VARS
	pwd, err := os.Getwd()
	envFile := path.Join(pwd, ".env")
	CheckErr(err)
	err = godotenv.Load(envFile)
	CheckErr(err)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = http.ListenAndServe(":"+port, router)
	CheckErr(err)
}
