package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lazyspell/Ecommerce_Backend/models"
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
		return "", fmt.Errorf("error in createToken when signing token %w", err)
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
		return nil, fmt.Errorf("error in parseToken")
	}
	if !verifiedToken.Valid {
		return nil, fmt.Errorf("error in parseToken")
	}

	return verifiedToken.Claims.(*UserClaims), nil
}

func jwtToken(user models.Users) (string, error) {
	tokenAuth := jwtauth.New("HS256", key, nil)
	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"first_name": user.FirstName})
	if err != nil {
		log.Println("issue generating the token")
	}

	return tokenString, nil

}
