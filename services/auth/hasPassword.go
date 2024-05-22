package authService

import "golang.org/x/crypto/bcrypt"


func HasPassword(password string) (string, error){
	has,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(has), nil
}


func ComparePassword(hasPw, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasPw), []byte(pwd))
}