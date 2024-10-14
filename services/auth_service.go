package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"golang.org/x/crypto/bcrypt"
)

func IntToBool(factor int32) bool {
	return factor != 0
}
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
func HashGID(input string) string {

	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 5)
	PanicOnError(err)
	return string(bytes)
}

func CompareGID(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func (s *Services) CreateToken(email string) (string, error) {
	id, err := s.Db.GetIdByEmail(s.Ctx, email)
	if err != nil {
		return "", err
	}

	jwtBuilder := jwt.NewBuilder()
	jwtBuilder.Issuer("gynasetu-server")
	jwtBuilder.IssuedAt(time.Now())
	jwtBuilder.JwtID(uuid.New().String())
	jwtBuilder.Claim("uid", int32(id))

	token, err := jwtBuilder.Build()
	if err != nil {
		return "", err
	}

	signedToken, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, s.Secret))
	if err != nil {
		return "", err
	}
	return string(signedToken), nil
}
func (s *Services) GetClaimsFromToken(token string) (jwt.Token, error) {
	parsedToken, err := jwt.Parse([]byte(token), jwt.WithKey(jwa.HS256, s.Secret))
	if err != nil {
		return nil, err
	}
	return parsedToken, nil
}
func (s *Services) GetIdFromToken(token string) (int32, error) {
	claims, err := s.GetClaimsFromToken(token)
	if err != nil {
		return 0, err
	}
	id, ok := claims.Get("uid")
	if !ok {
		return 0, fmt.Errorf("Token Invalid: Cannot Get The Required uid Claim")
	}
	return int32(id.(float64)), nil
}
