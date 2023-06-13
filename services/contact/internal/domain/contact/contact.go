package contact

import (

	"time"
	"github.com/MeibisuX673/practiceGo/pkg/type/email"

)

type Contact struct{

	ID int
	FirstName string
	LastName string
	Email email.Email
	cretedAt time.Time
	updateAt time.Time

}

func New(
	firstName, lastName string,
	email email.Email,
) (*Contact, error){

	contact := &Contact{
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		cretedAt: time.Now(),
		updateAt: time.Now(),
	}

	return contact, nil

}

