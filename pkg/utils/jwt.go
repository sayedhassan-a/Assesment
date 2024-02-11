package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type Claims struct {
	Type			string
	Email			string
	StandardClaims	jwt.StandardClaims
}

func (claims *Claims) Valid()error{
	err := claims.StandardClaims.Valid()
	if err != nil {
		return err
	}

	return nil
}

func getSecretKey()[]byte{
	secretKey := os.Getenv("SECRET_KEY")
	return []byte(secretKey)
}
func GenerateToken(email string) (accessToken string, refreshToken string, err error) {
	claims := &Claims{
		Type: "access",
		Email: email,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour).Unix(),
		},
	}

	refreshClaims := &Claims{
		Type: "refresh",
		Email: email,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 24 * 3).Unix(),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(getSecretKey())
	if err != nil {
		return "","", err
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(getSecretKey())
	if err != nil {
		return "","", err
	}

	return accessToken, refreshToken, nil
}

func ExtractClaims(tokenString string) (*Claims,error){
	token, err := jwt.ParseWithClaims(tokenString,&Claims{},func(token *jwt.Token)(interface{},error){
		return []byte(getSecretKey()), nil
	})

	if err != nil {
		return nil,err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil,fmt.Errorf("invalid token")
	}

	err = claims.Valid()
	if err != nil{
		return nil, err
	}

	return claims,nil
}

func Refresh(tokenString string) (string, string, error) {
	claims, err := ExtractClaims(tokenString)

	if err != nil {
		return "","",err
	}
	if claims.Type != "refresh" {
		return "","",fmt.Errorf("Invalid Token")
	}

	access,refresh,err:=GenerateToken(claims.Email)
	if err != nil {
		return "","",err
	}
	return access, refresh, nil
}