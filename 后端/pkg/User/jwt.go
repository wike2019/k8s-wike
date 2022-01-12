package User

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(Email ,Role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := User{
		Email: Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    Email,
		},
		Role: Role,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("wiki_boy_happy"))
	return token, err
}

func ParseToken(token string) (*User, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("wiki_boy_happy"), nil
	})
	if err != nil {
		return nil, err
	}

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*User); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}