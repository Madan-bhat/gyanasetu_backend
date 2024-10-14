package models

import "github.com/golang-jwt/jwt/v5"

type RegisterDTO struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	GID   string `json:"gid" validate:"required"`
}

type SignupDTO struct {
	Name     string `json:"name" validate:"required"`
	Phno     string `json:"phno" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type SelectRoleDTO struct {
	Role string `json:"role" validate:"required"`
}

type UserClaims struct {
	jwt.RegisteredClaims
	ID int32 `json:"id"`
}

type JWTSubject struct {
	Id   int32  `json:"id"`
	Name string `json:"id"`
}
