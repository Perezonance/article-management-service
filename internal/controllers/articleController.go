package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/Perezonance/article-management-service/internal/models"
	"github.com/Perezonance/article-management-service/internal/server"
	"github.com/Perezonance/article-management-service/internal/storage"
	log "github.com/Perezonance/article-management-service/internal/util/logger"
)

//Controller handles and processes requests and responses to the server as well as input validation
type Controller struct {
	s *server.Server
}

//NewController creates a controller for handling and processing requests and responses to the server
func NewController(s *server.Server) *Controller {
	return &Controller{s: s}
}

//GetArticlesHandler processes request and calls server to fetch all articles
//GET /articles
func (c *Controller) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	log.InfoLog("Request recieved: returning all articles.")

	arts, err := c.s.GetArticles()
	if err != nil {
		log.ErrorLog("Error while retrieving articles", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	res, err := json.Marshal(arts)
	if err != nil {
		log.ErrorLog("Error while marshaling response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusOK, string(res), w)
}

//PostArticleHandler processes request and calls server to create a new article
//POST /articles
func (c *Controller) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	var a models.NewArticle

	log.InfoLog("Request recieved: creating new article.")

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.ErrorLog("Error while decoding request payload", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	log.InfoLog(fmt.Sprintf("decoded request payload:\n%v", a))

	aID, err := c.s.CreateArticle(a)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while creating new article: payload\n%v\n", a), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	//TODO: return resource location
	writeRes(http.StatusCreated, fmt.Sprintf("%v", aID), w)
}

//GetArticleByIDHandler processes request and makes server call to fetch an article with given artID
//GET /articles/{articleID}
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	artID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/articles/"))
	if err != nil {
		log.ErrorLog("Error while parsing path URL", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	log.InfoLog(fmt.Sprintf("Request received: retrieving article with id%v", artID))

	art, err := c.s.GetArticleByID(artID)
	if err != nil {
		if err == storage.ErrResourceNotFound {
			log.ErrorLog(fmt.Sprintf("Error while retrieving article with id:%v", artID), err)
			writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
			return
		}
		log.ErrorLog(fmt.Sprintf("Error while retrieving article with id:%v", artID), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	res, err := json.Marshal(art)
	if err != nil {
		log.ErrorLog("Error while marshaling response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusOK, string(res), w)
}

//UpdateArticleByIDHandler processes request and makes server call to update an article with given artID
//PUT /articles/{articleID}
func (c *Controller) UpdateArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	artID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/articles/"))
	if err != nil {
		log.ErrorLog("Error while parsing path URL", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	var a models.Article

	log.InfoLog(fmt.Sprintf("Request received: updating article with id%v\n", artID))

	err = json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.ErrorLog("Error while decoding request payload", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	err = c.s.UpdateArticle(a)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while updating article with id:%v", artID), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	art, err := c.s.GetArticleByID(artID)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while returning article with id:%v", artID), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	res, err := json.Marshal(art)
	if err != nil {
		log.ErrorLog("Error while marshaling response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusAccepted, string(res), w)
}

//DeleteArticleByIDHandler processes request and makes server call to delete an article with given artID
//DELETE /articles/{articleID}
func (c *Controller) DeleteArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	artID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/articles/"))
	if err != nil {
		log.ErrorLog("Error while parsing path URL", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	log.InfoLog(fmt.Sprintf("Request received: deleting article with id%v\n", artID))

	err = c.s.DeleteArticle(artID)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while deleting article with id:%v", artID), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusOK, http.StatusText(http.StatusOK), w)
}

//GetArticleByUserIDHandler processes request and makes server call to fetch an article filtered
//with given userID
//GET /articles/{userID}
func (c *Controller) GetArticleByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/articles/id:"))
	if err != nil {
		log.ErrorLog("Error while parsing path URL", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	log.InfoLog(fmt.Sprintf("Request received: retrieving article with user id%v\n", userID))

	arts, err := c.s.GetArticlesByUser(userID)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while retrieving articles with user id:%v", userID), err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	res, err := json.Marshal(arts)
	if err != nil {
		log.ErrorLog("Error while marshaling response", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	writeRes(http.StatusOK, string(res), w)
}

func writeRes(statusCode int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	res := []byte(message)
	_, err := w.Write(res)
	if err != nil {
		log.ErrorLog("Error while writing to ResponseWriter", err)
		return
	}
	return
}
