package app

import accesstoken "github.com/rogik/store-auth-api/src/domain/access-token"

func StartApplication() {
	accessTokenService := accesstoken.NewService()
}
