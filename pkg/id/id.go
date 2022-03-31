package id

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateAPIKey() string {
	return uuid.New().String()
}

func GenerateRequestID() string {
	return uuid.New().String()
}

func GenerateTeamID() string {
	return uuid.New().String()
}

func GenerateTeamMemberKey() string {
	return uuid.New().String()
}

func GenerateUserName() string {
	return "zcashrpc"
}

func GenerateSecurePassword() string {
	return "password"
}

func HashAndSaltPassword(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}
