package models

import (
	"errors"
	"fmt"
	"regexp"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	if u.Age < 0 {
		return errors.New("age cannot be negative")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailRegex, u.Email)

	if !emailMatch {
		return errors.New("invalid email format")
	}

	phoneRegex := `^\d{9}$`
	phoneMatch, _ := regexp.MatchString(phoneRegex, fmt.Sprintf("%d", u.Phone))
	if !phoneMatch {
		return errors.New("invalid phone number format")
	}

	return nil
}
