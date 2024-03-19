package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// if expirationTimeHours != 24 {
	// 	t.Error("Expiration time should be 24 hours.")
	// }
	assert.EqualValues(t, expirationTimeHours, "Expiration time should be 24 hours.")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "Brand new access token should not be expires.")
	assert.EqualValues(t, "", at.AccessToken, "New access token should not have token id.")
	assert.True(t, at.UserId == 0, "New access token should not have an accosiated user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "Empty access token should be expired by default.")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "Access token expires 3 hours from now, so should not be expired.")
}
