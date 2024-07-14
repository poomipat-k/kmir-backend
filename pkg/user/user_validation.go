package user

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func validatePassword(password string) error {
	if password == "" {
		return PasswordRequiredError{}
	}
	if len(password) < 8 {
		return PasswordTooShortError{}
	}
	if len(password) > 60 {
		return PasswordTooLongError{}
	}
	return nil
}

func validateLoginPayload(payload LoginRequest) (string, error) {
	if payload.Username == "" {
		return "username", UsernameRequiredError{}
	}

	err := validatePassword(payload.Password)
	if err != nil {
		return "password", err
	}
	return "", nil
}

func comparePassword(inputPassword string, userPassword string) error {
	splitStr := strings.Split(userPassword, "_")
	if len(splitStr) != 2 {
		return errors.New("username or password is invalid")
	}
	hash := splitStr[0]
	salt := splitStr[1]
	// user provided password + salt compare to hashed
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(strings.Join([]string{inputPassword, salt}, "")))
	if err != nil {
		return err
	}
	return nil
}
