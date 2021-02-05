package main

import (
	"log"
	"net/http"

	"github.com/Perezonance/article-management-service/internal/controllers"
	"github.com/Perezonance/article-management-service/internal/server"
	"github.com/Perezonance/article-management-service/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	db := storage.NewMockDynamo()

	s := server.NewServer(db)

	c := controllers.NewController(s)

	r.HandleFunc("/articles", c.GetArticlesHandler).Methods(http.MethodGet)
	r.HandleFunc("/articles", c.PostArticleHandler).Methods(http.MethodPost)

	r.HandleFunc("/articles/{articleID}", c.GetArticleByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/articles/{articleID}", c.UpdateArticleByIDHandler).Methods(http.MethodPut)
	r.HandleFunc("/articles/{articleID}", c.DeleteArticleByIDHandler).Methods(http.MethodDelete)

	r.HandleFunc("/articles/{userID}", c.GetArticleByUserIDHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
