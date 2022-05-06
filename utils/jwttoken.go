package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var password = "Elaine I will love you forever and always"
var key = []byte(password)

type UserClaims struct {
	jwt.StandardClaims
	SessionID int64
}

func (u *UserClaims) Valid() error {
	if !u.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if u.SessionID == 0 {
		return fmt.Errorf("invalid Session ID")
	}

	return nil
}

func CreateToken(c *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodES512, c)
	signedToken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Error in createToken when signing token %w", err)
	}
	return signedToken, nil

}

func ParseToken(signedToken string) (*UserClaims, error) {
	verifiedToken, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if token.Method.Alg() != jwt.SigningMethodES512.Alg() {
			return nil, fmt.Errorf("invalid signing algorithm")
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("Error in parseToken")
	}
	if !verifiedToken.Valid {
		return nil, fmt.Errorf("Error in parseToken")
	}

	return verifiedToken.Claims.(*UserClaims), nil
}
