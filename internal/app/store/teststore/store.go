package teststore

import (
	"github.com/gtmartem/go-http-rest-api/internal/app/model"
	"github.com/gtmartem/go-http-rest-api/internal/app/store"
)


// Store ...
type Store struct {
	userRepository 	*UserRepository
}


// New ...
func New() store.Store {
	return &Store{}
}


// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}




