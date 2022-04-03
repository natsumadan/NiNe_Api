package routes

import (
	"github.com/gorilla/mux"
	"ninApi/pkg/controllers"
)

var ArticleApiRoutes = func(router *mux.Router) {
	router.HandleFunc("/articles/", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", controllers.GetArticle).Methods("GET")
	//	router.HandleFunc("/tags/{tagName}/{date}", controllers.GetTags).Methods("GET")

}
