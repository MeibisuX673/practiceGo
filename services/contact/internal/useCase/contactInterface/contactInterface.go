package contactinterface

import(

	"github.com/MeibisuX673/practiceGo/services/contact/internal/domain/contact"

)

type ContactInterface interface{

	Create(*contact.Contact) *contact.Contact
	GetCollection() []*contact.Contact
	GetById(*contact.Contact) *contact.Contact
	Update(contact.Contact) *contact.Contact
	Delete(id int)

}