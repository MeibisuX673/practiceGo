package storage

import "github.com/MeibisuX673/practiceGo/services/contact/internal/domain/contact"


type Storage interface {

	ContactRepositoryInterface
	
}

type ContactRepositoryInterface interface{

	Create(*contact.Contact) *contact.Contact
	Update(id int, updateFn func(c *contact.Contact) *contact.Contact) *contact.Contact
	GetCollection() []*contact.Contact
	GetById(id int) *contact.Contact
	Delete(id int)
	
}