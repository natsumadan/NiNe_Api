package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"ninApi/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.ArticleApiRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
