package accesstoken

import (
	"github.com/rogik/store-auth-api/src/utils/errors"
)

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type Service interface {
	GetById(string) (*AccessToken, *errors.RestError)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) GetById(string) (*AccessToken, *errors.RestError) {
	return nil, nil
}
