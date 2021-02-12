package server

import (
	"fmt"

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
	arts := make([]models.Article, len(ids))
	quit := make(chan bool)
	errc := make(chan error)
	done := make(chan error)

	for i, v := range ids {
		go func(i int, v int) {
			art, err := s.GetArticleByID(v)
			ch := done
			arts[i] = art

			if err != nil {
				log.ErrorLog(fmt.Sprintf("Thread %v returned error", i), err)
				ch = errc
			}
			select {
			case ch <- err:
				return
			case <-quit:
				return
			}
		}(i, v)
	}
	count := 0
	blank := make([]models.Article, 0)

	for {
		select {
		case err := <-errc:
			close(quit)
			return blank, err
		case <-done:
			count++
			if count == len(arts) {
				if len(arts) != len(ids) {
					return blank, fmt.Errorf("uneven ids input to articles output")
				}
				return arts, nil
			}
		}
	}

	// var (
	// 	mu    = &sync.Mutex{}
	// 	arts  = make([]models.Article, len(ids))
	// 	echan = make(chan error)
	// )

	// log.DebugLog(fmt.Sprintf("Number of ids:%v", len(ids)))

	// var wg sync.WaitGroup
	// for i, v := range ids {
	// 	log.DebugLog(fmt.Sprintf("Thread %v created", i))
	// 	wg.Add(1)
	// 	go func(i int, v int) {
	// 		defer log.DebugLog(fmt.Sprintf("Thread %v completed", i))
	// 		defer wg.Done()
	// 		mu.Lock()
	// 		art, err := s.GetArticleByID(v)
	// 		mu.Unlock()
	// 		if err != nil {
	// 			log.ErrorLog(fmt.Sprintf("Error while thread %v was requesting article from db with id:%v", i, v), err)
	// 		}
	// 		echan <- err
	// 		arts = append(arts, art)
	// 	}(i, v)
	// }

	// log.DebugLog("fetched all ids, processing error channel")
	// select {
	// case err <- echan:
	// 	log.ErrorLog("error recieved through channel", err)
	// default:
	// 	log.DebugLog("error not parsed through error channel")
	// }
	// for i := 0; i < len(arts); i++ {
	// 	err := <-echan
	// 	if err != nil {
	// 		log.DebugLog("Non nil value recieved through error channel")
	// 		blank := make([]models.Article, 0)
	// 		return blank, err
	// 	}
	// }
	// log.DebugLog("fetched all ids, error channel processed")

	// wg.Wait()
	// log.DebugLog(fmt.Sprintf("returning fetched articles:\n%v", arts))

	// return arts, nil
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

//CreateArticles creates a new article given the article data model and returns the newly issued ID
//POST /articles
func (s *Server) CreateArticles(arts []models.NewArticle) ([]int, error) {
	ids := make([]int, len(arts))
	quit := make(chan bool)
	errc := make(chan error)
	done := make(chan error)

	for i, v := range arts {
		go func(i int, v models.NewArticle) {
			id, err := s.CreateArticle(v)
			ch := done
			ids[i] = id

			if err != nil {
				log.ErrorLog(fmt.Sprintf("Thread %v returned error", i), err)
				ch = errc
			}
			select {
			case ch <- err:
				return
			case <-quit:
				return
			}
		}(i, v)
	}
	count := 0
	blank := make([]int, 0)

	for {
		select {
		case err := <-errc:
			close(quit)
			return blank, err
		case <-done:
			count++
			if count == len(arts) {
				if len(arts) != len(ids) {
					return blank, fmt.Errorf("uneven newArticle input to article ids output")
				}
				return ids, nil
			}
		}
	}

	// var (
	// 	mu     = &sync.Mutex{}
	// 	artIDs []int
	// 	echan  = make(chan error, len(arts))
	// )

	// log.DebugLog("creating multiple articles")

	// var wg sync.WaitGroup
	// for i, v := range arts {
	// 	wg.Add(1)
	// 	go func(i int, v models.NewArticle) {
	// 		log.DebugLog(fmt.Sprintf("Thread %v created..", i))
	// 		defer wg.Done()
	// 		mu.Lock()
	// 		id, err := s.CreateArticle(arts[i])
	// 		mu.Unlock()
	// 		if err != nil {
	// 			log.ErrorLog(fmt.Sprintf("Error while creating new article:\n%v", v), err)
	// 		}
	// 		echan <- err
	// 		artIDs = append(artIDs, id)
	// 	}(i, v)
	// }
	// //Check Error channel for any errors
	// for i := 0; i < len(arts); i++ {
	// 	err := <-echan
	// 	if err != nil {
	// 		log.DebugLog("Non nil value recieved through error channel")
	// 		blank := make([]int, 0)
	// 		return blank, err
	// 	}
	// }
	// wg.Wait()

	// log.DebugLog("articles created, processing error channel")

	// log.DebugLog("returning generated article IDs")

	// return artIDs, nil
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
