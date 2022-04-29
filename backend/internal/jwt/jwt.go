package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

const (
	tokenLifetime  = 2 * time.Hour
	minutesToRenew = 5
)

var (
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidTokenClaims   = jwt.ErrTokenInvalidClaims
)

// Create generate new token with given credentials and secret.
func Create(login, secret string) (string, time.Time, error) {
	expireTime := time.Now().Add(tokenLifetime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	})

	signedToken, err := token.SignedString([]byte(secret))

	return signedToken, expireTime, err
}

// Parse token and return login and expiration time.
func Parse(token, secret string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &Claims{}, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidSigningMethod
		}

		return []byte(secret), nil
	})
}

// NeedUpdate returns true if token need to be updated.
func NeedUpdate(token *jwt.Token) (bool, error) {
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return false, jwt.ErrTokenInvalidClaims
	}

	// Get expire time from claims
	expireTime := claims.ExpiresAt.Time
	nowTime := time.Now()
	timeDiff := expireTime.Sub(nowTime)

	// Check that token needs to be updated
	if timeDiff.Minutes() > minutesToRenew {
		return false, nil
	}

	return true, nil
}
