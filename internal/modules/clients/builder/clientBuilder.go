package builder

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/dto"
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/entities"
	"gopkg.in/guregu/null.v4"
)

type ClientBuilder struct {
	client entities.Client
}

func NewClientBuild() *ClientBuilder {
	return &ClientBuilder{}
}

func (b *ClientBuilder) CreateClient(clientDto dto.Client) *ClientBuilder {
	b.client = entities.Client{
		FullName:      clientDto.FullName,
		Gender:        null.NewString(clientDto.Gender, true),
		Type:          null.NewString(clientDto.Type, true),
		BirthDate:     null.NewString(clientDto.BirthDate, true),
		PlaceOfBirth:  null.NewString(clientDto.PlaceOfBirth, true),
		Nationality:   null.NewString(clientDto.Nationality, true),
		MaritalStatus: null.NewString(clientDto.MaritalStatus, true),
	}

	return b
}

func (b *ClientBuilder) WithAddress(addreses []dto.Address) *ClientBuilder {
	for _, address := range addreses {
		a := entities.Address{
			Street:   null.NewString(address.Street, true),
			District: null.NewString(address.District, true),
			City:     null.NewString(address.City, true),
			ZipCode:  null.NewString(address.ZipCode, true),
			State:    null.NewString(address.State, true),
			Type:     null.NewString(address.Type, true),
		}
		b.client.AddAddress(a)
	}

	return b
}

func (b *ClientBuilder) WithContacts(contacts []dto.Contact) *ClientBuilder {
	for _, contact := range contacts {
		c := entities.Contact{
			Phone:          null.NewString(contact.Phone, true),
			MobileOperator: null.NewString(contact.MobileOperator, true),
			Site:           null.NewString(contact.Site, true),
			Type:           null.NewString(contact.Type, true),
		}
		b.client.AddContact(c)
	}

	return b
}

func (b *ClientBuilder) WithDocuments(documents dto.Documents) *ClientBuilder {
	d := entities.Documents{
		Rg:            null.NewString(documents.Rg, true),
		CPForCNPJ:     null.NewString(documents.CPForCNPJ, true),
		TituloEleitor: null.NewString(documents.TituloEleitor, true),
		CTPS:          null.NewString(documents.CTPS, true),
		PIS:           null.NewString(documents.PIS, true),
		CNH:           null.NewString(documents.CNH, true),
		Passport:      null.NewString(documents.Passport, true),
		Reservista:    null.NewString(documents.Reservista, true),
		IE:            null.NewString(documents.IE, true),
		IM:            null.NewString(documents.IM, true),
	}
	b.client.AddDocuments(d)

	return b
}

func (b *ClientBuilder) WithFiliation(filiation dto.Filiation) *ClientBuilder {
	f := entities.Filiation{
		MotherName: null.NewString(filiation.MotherName, true),
		FatherName: null.NewString(filiation.FatherName, true),
	}
	b.client.AddFiliation(f)

	return b
}

func (b *ClientBuilder) WithBankAccounts(accounts []dto.BankAccount) *ClientBuilder {
	for _, account := range accounts {
		a := entities.BankAccount{
			Type:    null.NewString(account.Type, true),
			Name:    null.NewString(account.Name, true),
			Agency:  null.NewString(account.Agency, true),
			Account: null.NewString(account.Account, true),
			Pix:     null.NewString(account.Pix, true),
		}
		b.client.AddBankAccount(a)
	}

	return b
}

func (b *ClientBuilder) Get() *entities.Client {
	return &b.client
}

func (b *ClientBuilder) Address(address dto.Address) *entities.Address {
	return &entities.Address{
		Street:   null.NewString(address.Street, true),
		District: null.NewString(address.District, true),
		City:     null.NewString(address.City, true),
		ZipCode:  null.NewString(address.ZipCode, true),
		State:    null.NewString(address.State, true),
		Type:     null.NewString(address.Type, true),
	}
}

func (b *ClientBuilder) Contact(contact dto.Contact) *entities.Contact {
	return &entities.Contact{
		Phone:          null.NewString(contact.Phone, true),
		MobileOperator: null.NewString(contact.MobileOperator, true),
		Site:           null.NewString(contact.Site, true),
		Type:           null.NewString(contact.Type, true),
	}
}

func (b *ClientBuilder) BankAccount(account dto.BankAccount) *entities.BankAccount {
	return &entities.BankAccount{
		Type:    null.NewString(account.Type, true),
		Name:    null.NewString(account.Name, true),
		Agency:  null.NewString(account.Agency, true),
		Account: null.NewString(account.Account, true),
		Pix:     null.NewString(account.Pix, true),
	}
}

func (b *ClientBuilder) Documents(documents dto.Documents) *entities.Documents {
	return &entities.Documents{
		Rg:            null.NewString(documents.Rg, true),
		CPForCNPJ:     null.NewString(documents.CPForCNPJ, true),
		TituloEleitor: null.NewString(documents.TituloEleitor, true),
		CTPS:          null.NewString(documents.CTPS, true),
		PIS:           null.NewString(documents.PIS, true),
		CNH:           null.NewString(documents.CNH, true),
		Passport:      null.NewString(documents.Passport, true),
		Reservista:    null.NewString(documents.Reservista, true),
		IE:            null.NewString(documents.IE, true),
		IM:            null.NewString(documents.IM, true),
	}
}

func (b *ClientBuilder) Filiation(filiation dto.Filiation) *entities.Filiation {
	return &entities.Filiation{
		MotherName: null.NewString(filiation.MotherName, true),
		FatherName: null.NewString(filiation.FatherName, true),
	}
}
