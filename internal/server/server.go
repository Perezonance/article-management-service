package server

import "github.com/Perezonance/article-management-service/internal/storage"

//Server processes the data models and handles business logic for the server
type Server struct {
	db *storage.Storage
}

//NewServer creates a server and returns it given a storage
func NewServer(d *storage.Storage) *Server {
	return &Server{db: d}
}
