package storage

import "github.com/Perezonance/article-management-service/internal/models"

var (
	idCounter = 1
)

//MockDynamo emulates a key value store db with an Articles table
type MockDynamo struct {
	ArticlesTable map[int]models.Article
}

//NewMockDynamo creates a new MockDynamo DB with a blank Articles table
func NewMockDynamo() *MockDynamo {
	return &MockDynamo{ArticlesTable: make(map[int]models.Article)}
}

//GetArticleByID returns an article given an id
func (mdb *MockDynamo) GetArticleByID(id int) (models.Article, error) {
	article := mdb.ArticlesTable[id]
	//Check to see if key contains non-zero value
	if article == (models.Article{}) {
		return models.Article{}, ErrResourceNotFound
	}
	return article, nil
}

//GetAllArticles returns all articles in the in-memory db
func (mdb *MockDynamo) GetAllArticles() ([]models.Article, error) {
	var articles []models.Article
	for _, v := range mdb.ArticlesTable {
		articles = append(articles, v)
	}
	return articles, nil
}

//GetArticleByUserID returns all articles filtered by a particular userId
func (mdb *MockDynamo) GetArticleByUserID(userID int) ([]models.Article, error) {
	var articles []models.Article
	for _, v := range mdb.ArticlesTable {
		if v.UserID == userID {
			articles = append(articles, v)
		}
	}
	return articles, nil
}

//CreateArticle adds a new article into the in-mem mock db
func (mdb *MockDynamo) CreateArticle(art models.NewArticle) (int, error) {
	//TODO: Needs cleaner solution - work around memory counter and json decoding
	var insertArt = models.Article{
		ArticleID: idCounter,
		UserID:    art.UserID,
		Title:     art.Title,
		Body:      art.Body,
	}
	idCounter++
	mdb.ArticlesTable[insertArt.ArticleID] = insertArt
	return insertArt.ArticleID, nil

}

//UpdateArticle replaces an existing article with a new one
func (mdb *MockDynamo) UpdateArticle(id int, article models.Article) error {
	_, err := mdb.GetArticleByID(id)
	if err != nil {
		return err
	}
	mdb.ArticlesTable[id] = article
	return nil
}

//DeleteArticle removes an article from the in-memory mock db
func (mdb *MockDynamo) DeleteArticle(id int) error {
	_, err := mdb.GetArticleByID(id)
	if err != nil {
		return err
	}
	delete(mdb.ArticlesTable, id)
	return nil
}
