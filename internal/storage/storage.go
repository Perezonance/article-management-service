package storage

import "github.com/Perezonance/article-management-service/internal/models"

//Storage defines the behavior for a db accessing tool
type Storage interface {
	GetArticleByID(int) (models.Article, error)
	GetAllArticles() ([]models.Article, error)
	GetArticleByUserID(int) ([]models.Article, error)
	CreateArticle(models.NewArticle) (int, error)
	UpdateArticle(int, models.Article) error
	DeleteArticle(int) error
}
