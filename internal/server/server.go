package server

import (
	"fmt"
	"sync"

	"github.com/Perezonance/article-management-service/internal/models"
	"github.com/Perezonance/article-management-service/internal/storage"
	log "github.com/Perezonance/article-management-service/internal/util/logger"
)

//Server processes the data models and handles business logic for the server
type Server struct {
	db storage.Storage
}

//NewServer creates a server and returns it given a storage
func NewServer(d storage.Storage) *Server {
	return &Server{db: d}
}

//GetArticles returns all the articles in the db
//GET /articles
func (s *Server) GetArticles() ([]models.Article, error) {
	articles, err := s.db.GetAllArticles()
	if err != nil {

	}
	return articles, nil
}

//GetArticleByID returns the article represented by the articleId given
//GET /articles/{articleId}
func (s *Server) GetArticleByID(id int) (models.Article, error) {
	var (
		article models.Article
	)
	article, err := s.db.GetArticleByID(id)
	if err != nil {
		log.ErrorLog(fmt.Sprintf("Error while requesting article from db with id:%v", id), err)
		return article, err
	}
	return article, nil
}

//GetArticlesByIDs returns all articles requested given a slice of ids
//GET /articles?ids=id1,id2,id3,idn...
func (s *Server) GetArticlesByIDs(ids []int) ([]models.Article, error) {
	var (
		mu   = &sync.Mutex{}
		arts = make([]models.Article, 0)
	)

	var wg sync.WaitGroup
	for i := range ids {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			art, err := s.GetArticleByID(i)
			if err != nil {
				//There needs to be error handling for this situation during the mutex lock
				log.ErrorLog(fmt.Sprintf("Error while requesting article from db with id:%v", i), err)
			}
			arts = append(arts, art)
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	return arts, nil
}

//CreateArticle creates a new article given the article data model and returns the newly issued ID
//POST /articles
func (s *Server) CreateArticle(a models.NewArticle) (int, error) {
	id, err := s.db.CreateArticle(a)
	if err != nil {
		log.ErrorLog("Error while creating new log", err)
		return 0, err
	}
	return id, nil
}

//UpdateArticle updates an existing article with the given data model and id
//PUT /articles/{articleId}
func (s *Server) UpdateArticle(a models.Article) error {
	err := s.db.UpdateArticle(a.ArticleID, a)
	if err != nil {
		//TODO: Check for 404
		log.ErrorLog(fmt.Sprintf("Error while updating log with id:%v", a.UserID), err)
		return err
	}
	return nil
}

//DeleteArticle deletes an article given the id
//DELETE /articles/{articleId}
func (s *Server) DeleteArticle(id int) error {
	err := s.db.DeleteArticle(id)
	if err != nil {
		//TODO: Check for 404
		log.ErrorLog(fmt.Sprintf("Error while deleting log with id:%v", id), err)
		return err
	}
	return nil
}

//GetArticlesByUser returns a list of all articles written by the given user
//GET /articles/user/{userId}
func (s *Server) GetArticlesByUser(userID int) ([]models.Article, error) {
	var arts []models.Article
	arts, err := s.db.GetArticleByUserID(userID)
	if err != nil {
		//TODO: Check for 404
		log.ErrorLog(fmt.Sprintf("Error while fetching articles with user id:%v", userID), err)
		return arts, err
	}
	return arts, nil
}
