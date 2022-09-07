package utils

import "golang.org/x/crypto/bcrypt"

func GenPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func ComparePwd(pwd1 string, pwd2 string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}
