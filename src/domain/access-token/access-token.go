package accesstoken

import (
	"time"
)

const (
	expirationTimeHours = 12
)

type AccessToken struct {
	AccessToken string `json:"accessToken`
	UserId      int64  `json:"userId`
	ClientId    int64  `json:"clientId`
	Expires     int64  `json:"expires`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTimeHours * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}
