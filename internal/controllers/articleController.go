package controllers

import (
	"net/http"

	"github.com/Perezonance/article-management-service/internal/server"
)

//Controller handles and processes requests and responses to the server as well as input validation
type Controller struct {
	s *server.Server
}

//NewController creates a controller for handling and processing requests and responses to the server
func NewController(s *server.Server) *Controller {
	return &Controller{s: s}
}

func (c *Controller) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) PostArticleHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) UpdateArticleByIDHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) DeleteArticleByIDHandler(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetArticleByUserIDHandler(w http.ResponseWriter, r *http.Request) {

}
