package teststore

import (
	"github.com/gtmartem/go-http-rest-api/internal/app/model"
)


// Store ...
type Store struct {
	userRepository 	*UserRepository
}


// New ...
func New() *Store {
	return &Store{}
}


// User ...
func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
		users: make(map[string]*model.User),
	}
	return s.userRepository
}




