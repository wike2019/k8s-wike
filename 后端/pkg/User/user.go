package User

import "github.com/dgrijalva/jwt-go"

type EmailCheck struct {
	Email   string   `json:"email" binding:"required,email"`
}

type LoginUser struct {
	Email   string   `json:"email" binding:"required,email"`
	Code   string   `json:"code" binding:"required"`
}
type User struct {
	Email string
	jwt.StandardClaims
	Role string
}

