package storage

import "errors"

var (
	//ErrResourceNotFound is thrown when the db cannot return the resource requested
	ErrResourceNotFound = errors.New("resource requested was not found")
)
