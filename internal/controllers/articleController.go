package controllers

import "net/http"

//Controller handles and processes requests and responses to the server as well as input validation
type Controller struct {
}

//NewController creates a controller for handling and processing requests and responses to the server
func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) GetArticles(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) PostArticle(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetArticleByID(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) UpdateArticleByID(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) DeleteArticleByID(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetArticleByUserID(w http.ResponseWriter, r *http.Request) {

}
