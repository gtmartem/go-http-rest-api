package teststore_test

import (
	"github.com/gtmartem/go-http-rest-api/internal/app/model"
	"github.com/gtmartem/go-http-rest-api/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)
	err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}


func TestUserRepository_FindByEmail(t *testing.T) {
	s := teststore.New()

	email := model.TestUser(t).Email
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}

