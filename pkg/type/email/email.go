package email

import (

	"errors"
	"regexp"
	"strings"
	
)

type Email struct{
	value string
}

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func New(email string) (Email, error){

	if len(email) == 0 {
		return Email{}, nil
	}

	if !regexpEmail.MatchString(strings.ToLower(email)) {
		return Email{}, errors.New("invalid email format")
	}

	return Email{value: email}, nil

}

func (e Email) String() string{
	return e.value
}

