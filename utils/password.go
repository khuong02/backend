package utils

import (
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type ValidationError struct {
	Message string
	Field   string
	Tag     string
}

func (ve *ValidationError) Error() string {
	return ve.Message
}

func Password(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return &ValidationError{
			Message: "Password should be of 8 characters long",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}

	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one lower case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}

	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one upper case character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}

	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one digit",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}

	if !done {
		return &ValidationError{
			Message: "Password should contain atleast one special character",
			Field:   "password",
			Tag:     "strong_password",
		}
	}

	return nil
}
