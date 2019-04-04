package common

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	TokenID  string `json:"tokenId"`
	ExpireAt int64  `json:"exp"`

	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

func (c *Claims) Valid() error {
	if c.ExpireAt <= time.Now().UTC().Unix() {
		return errors.New("Expired all already")
	}
	return nil
}

func GenToken(secret []byte, c *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(secret)
}

func ParseToken(secret []byte, token string) (*Claims, error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	tokenObj, err := jwt.ParseWithClaims(token, &Claims{}, func(tokenObj *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := tokenObj.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tokenObj.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokenObj.Claims.(*Claims)
	if !ok {
		return nil, errors.New("parse token fail")
	}
	if tokenObj.Valid {
		return claims, nil
	} else {
		return nil, errors.New("token is invalid")
	}
}
