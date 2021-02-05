package main

import (
	"log"
	"net/http"

	"github.com/Perezonance/article-management-service/internal/controllers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	c := controllers.NewController()

	r.HandleFunc("/articles", c.GetArticles).Methods(http.MethodGet)
	r.HandleFunc("/articles", c.PostArticle).Methods(http.MethodPost)

	r.HandleFunc("/articles/{articleID}", c.GetArticleByID).Methods(http.MethodGet)
	r.HandleFunc("/articles/{articleID}", c.UpdateArticleByID).Methods(http.MethodPut)
	r.HandleFunc("/articles/{articleID}", c.DeleteArticleByID).Methods(http.MethodDelete)

	r.HandleFunc("/articles/{userID}", articleByUserIDHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
