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
	"github.com/gorilla/mux"
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
//GET /articles?ids=1,3,127, 13048203
func (c *Controller) GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ids    = strings.Split(r.URL.Query().Get("ids"), ",")
		arts   []models.Article
		intIDs = make([]int, len(ids))
	)

	log.DebugLog(fmt.Sprintf("Number of Ids requested:%v", len(ids)))
	if len(ids) > 1 {
		log.InfoLog("Request recieved: returning all articles.")
		a, err := c.s.GetArticles()
		if err != nil {
			log.ErrorLog("Error while retrieving articles", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
		log.DebugLog(fmt.Sprintf("Request processing: retrieved all articles"))
		arts = a
	} else {
		log.InfoLog(fmt.Sprintf("Request recieved: returning articles for ids:%v", intIDs))
		for i, s := range ids {
			intIDs[i], _ = strconv.Atoi(s)
		}
		a, err := c.s.GetArticlesByIDs(intIDs)
		if err != nil {
			log.ErrorLog("Error while retrieving articles", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
		arts = a
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

	a := make([]models.NewArticle, 10)

	log.InfoLog("Request recieved: creating new article(s).")

	err := json.NewDecoder(r.Body).Decode(&a)
	if err != nil {
		log.ErrorLog("Error while decoding request payload", err)
		writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	log.InfoLog(fmt.Sprintf("decoded request payload:\n%v", a))

	if len(a) > 1 {
		aIDs, err := c.s.CreateArticles(a)
		if err != nil {
			log.ErrorLog(fmt.Sprintf("Error while creating new article: payload\n%v\n", a), err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}

		res, err := json.Marshal(aIDs)
		if err != nil {
			log.ErrorLog("Error while marshaling response", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
		writeRes(http.StatusAccepted, string(res), w)
	} else {
		aID, err := c.s.CreateArticle(a[0])
		if err != nil {
			log.ErrorLog(fmt.Sprintf("Error while creating new article: payload\n%v\n", a), err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}

		res, err := json.Marshal(aID)
		if err != nil {
			log.ErrorLog("Error while marshaling response", err)
			writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
			return
		}
		writeRes(http.StatusAccepted, string(res), w)
	}
	return
}

//GetMultArticleByIDHandler processes request and makes server call to fetch an article with given artID
//GET /articles/{articleID}
func (c *Controller) GetMultArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("ids")

	log.InfoLog(ids)

	//log.InfoLog(fmt.Sprintf("Request received: retrieving article with id%v", artID))

	// art, err := c.s.GetArticleByID(artID)
	// if err != nil {
	// 	if err == storage.ErrResourceNotFound {
	// 		log.ErrorLog(fmt.Sprintf("Error while retrieving article with id:%v", artID), err)
	// 		writeRes(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
	// 		return
	// 	}
	// 	log.ErrorLog(fmt.Sprintf("Error while retrieving article with id:%v", artID), err)
	// 	writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
	// 	return
	// }

	// res, err := json.Marshal(art)
	// if err != nil {
	// 	log.ErrorLog("Error while marshaling response", err)
	// 	writeRes(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
	// 	return
	// }
	// writeRes(http.StatusOK, string(res), w)
}

//GetArticleByIDHandler processes request and makes server call to fetch an article with given artID
//GET /articles/{articleID}
func (c *Controller) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	artID, err := strconv.Atoi(params["articleID"])
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
	params := mux.Vars(r)
	artID, err := strconv.Atoi(params["articleID"])
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
	params := mux.Vars(r)
	artID, err := strconv.Atoi(params["articleID"])
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
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["userID"])
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
