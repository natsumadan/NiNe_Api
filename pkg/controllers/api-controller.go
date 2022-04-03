package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"ninApi/pkg/models"
	"ninApi/pkg/utility"
	"strconv"
)

var NewArticle models.Article

func GetArticle(write http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	articleId := vars["articleId"]
	Id, err := strconv.ParseInt(articleId, 0, 0)
	if err != nil {
		fmt.Println("Parsing error")
	}
	articles, _ := models.GetArticle(Id)
	result, _ := json.Marshal(articles)
	write.Header().Set("Content", "Article/json")
	write.WriteHeader(http.StatusOK)
	write.Write(result)
}
func CreateArticle(write http.ResponseWriter, req *http.Request) {
	CreateArticle := &models.Article{}
	utility.ParseBody(req, CreateArticle)
	ar := CreateArticle.CreateArticle()
	result, _ := json.Marshal(ar)
	write.WriteHeader(http.StatusOK)
	write.Write(result)
}
