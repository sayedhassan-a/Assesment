package utils

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Hash(password string)string{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),10)
	if err != nil{
		log.Panic(err)
	}
	return string(bytes)
}

func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if err != nil {
		return false
	}
	return true
}