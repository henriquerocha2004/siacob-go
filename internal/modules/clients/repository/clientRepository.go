package repository

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/entities"
	"github.com/jmoiron/sqlx"
)

type ClientRepository interface {
	Create(client entities.Client) (int, error)
	Update(client entities.Client) error
	Delete(id int)
	FindAll() ([]entities.Client, error)
	FindOne(id int) (*entities.Client, error)
	CreateAddress(address entities.Address, transaction *sqlx.Tx) error
	UpdateAddress(address entities.Address) error
	FindAddressByClient(clientId int) (*[]entities.Address, error)
	DeleteAddress(id int)
	CreateContact(contact entities.Contact, transaction *sqlx.Tx) error
	UpdateContact(contact entities.Contact) error
	FindContactsByClient(clientId int) (*[]entities.Contact, error)
	DeleteContact(id int)
	CreateBankAccount(account entities.BankAccount, transaction *sqlx.Tx) error
	UpdateBankAccount(account entities.BankAccount) error
	FindBankAccountsByClient(clientId int) (*[]entities.BankAccount, error)
	DeleteBankAccount(id int)
	CreateDocuments(documents entities.Documents, transaction *sqlx.Tx) error
	UpdateDocuments(documents entities.Documents) error
	FindDocumentsByClient(clientId int) (entities.Documents, error)
	DeleteDocuments(id int)
	CreateFiliation(filiation entities.Filiation, transaction *sqlx.Tx) error
	UpdateFiliation(filiation entities.Filiation) error
	FindFiliationByClient(clientId int) (entities.Filiation, error)
	DeleteFiliation(id int)
}
