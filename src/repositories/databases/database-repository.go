package databases

import "github.com/rogik/store-auth-api/src/utils/errors"

type DbRepository interface {
	GetById(string) (*AccessToken, *errors.RestError)
}
