package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccessTokenConstants(t *testing.T) {
	// 	if expirationTime != 24 {
	// 		t.Error("expiration time should be 24 hours")
	// 	}
	assert.EqualValues(t, 24,expirationTime, "expiration time should be 24, but got %d",expirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be nil")
	// if at.IsExpired() {
	// 	t.Error("brand new access token should not be nil")
	// }

	assert.Empty(t, at.AccessToken, "new access token should not have define access token id")

	// if at.AccessToken != "" {
	// 	t.Error("new access token should not have define a")
	// }

	assert.Zero(t, at.UserId,"new access token should not have associated user id")
		// if at.UserId != 0 {
		// 	t.Error("new access token should not have associated user id")
		// }
}

func TestAccessTokenExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }

	at.Expires = time.Now().UTC().Add(3*time.Hour).Unix()

	assert.False(t, at.IsExpired(), "Access token expired three hour from now should not be expired")
	// if at.IsExpired() {
	// 	t.Error("Access token expired three hour from now should not be expired")
	// }
}