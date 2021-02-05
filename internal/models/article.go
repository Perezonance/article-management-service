package models

type (
	//Article provides the data model for an Article resource
	Article struct {
		UserID    int    `json:"userID"`
		ArticleID int    `json:"articleID"`
		Title     string `json:"title"`
		Body      string `json:"body"`
	}

	//NewArticle provices the data model for the request paylod of a new Article
	NewArticle struct {
		UserID int    `json:"userID"`
		Title  string `json:"title"`
		Body   string `json:"body"`
	}
)
