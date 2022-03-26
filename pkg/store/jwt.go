package store

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"sync"
	"time"
)

// a thread safe JWT Token Store with validation and expiration check
// use for caching
type JwtStore struct {
	keys sync.Map
}

func (s *JwtStore) GetValidJwt(key string) string {
	if token, exists := s.keys.Load(key); exists {
		tokenStr := token.(string)
		if isTokenValid(tokenStr) {
			return tokenStr
		}
	}

	return ""
}

func (s *JwtStore) StoreJwtToken(key, token string) {
	s.keys.Store(key, token)
}

// Valid JWT token is defined as one that is (1) parsable (2) not verified (3) has not expired
func isTokenValid(tokenStr string) bool {
	token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})
	if err != nil {
		// Something is corrupt with this token
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// Something is corrupt with this token
		return true
	}

	var tm time.Time
	switch iat := claims["exp"].(type) {
	case float64:
		tm = time.Unix(int64(iat), 0)
	case json.Number:
		v, _ := iat.Int64()
		tm = time.Unix(v, 0)
	}

	// If token does NOT expire in less than 1 hour or more, then it's valid
	return time.Now().Add(time.Hour).Unix() < tm.Unix()
}
