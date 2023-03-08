package entities

import (
	"encoding/json"

	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type Client struct {
	Id            int         `db:"id,omitempty"`
	FullName      string      `db:"full_name"`
	Gender        null.String `db:"gender"`
	Type          null.String `db:"type"`
	BirthDate     null.String `db:"birth_date"`
	PlaceOfBirth  null.String `db:"place_of_birth"`
	Nationality   null.String `db:"nationality"`
	MaritalStatus null.String `db:"marital_status"`
	timestamps.TimeStamps
	addresses    []Address
	documents    Documents
	contacts     []Contact
	bankAccounts []BankAccount
	filiation    Filiation
}

func Create() *Client {
	return &Client{}
}

func (c *Client) AddAddress(address Address) {
	c.addresses = append(c.addresses, address)
}

func (c *Client) GetAddresses() []Address {
	return c.addresses
}

func (c *Client) AddDocuments(documents Documents) {
	c.documents = documents
}

func (c *Client) GetDocuments() *Documents {
	return &c.documents
}

func (c *Client) AddFiliation(filiation Filiation) {
	c.filiation = filiation
}

func (c *Client) GetFiliation() *Filiation {
	return &c.filiation
}

func (c *Client) AddContact(contact Contact) {
	c.contacts = append(c.contacts, contact)
}

func (c *Client) GetContacts() []Contact {
	return c.contacts
}

func (c *Client) AddBankAccount(account BankAccount) {
	c.bankAccounts = append(c.bankAccounts, account)
}

func (c *Client) GetBankAccounts() []BankAccount {
	return c.bankAccounts
}

func (c *Client) MarshalJSON() ([]byte, error) {
	json, err := json.Marshal(struct {
		FullName      string        `json:"full_name"`
		Gender        null.String   `json:"gender"`
		Type          null.String   `json:"type"`
		BirthDate     null.String   `json:"birth_date"`
		PlaceOfBirth  null.String   `json:"place_of_birth"`
		Nationality   null.String   `json:"nationality"`
		MaritalStatus null.String   `json:"marital_status"`
		Address       []Address     `json:"addresses"`
		Documents     Documents     `json:"documents"`
		Contacts      []Contact     `json:"contacts"`
		BankAccounts  []BankAccount `json:"bank_accounts"`
		Filiation     Filiation     `json:"filiation"`
	}{
		FullName:      c.FullName,
		Gender:        c.Gender,
		Type:          c.Type,
		BirthDate:     c.BirthDate,
		PlaceOfBirth:  c.PlaceOfBirth,
		Nationality:   c.Nationality,
		MaritalStatus: c.MaritalStatus,
		Address:       c.addresses,
		Documents:     c.documents,
		Contacts:      c.contacts,
		BankAccounts:  c.bankAccounts,
		Filiation:     c.filiation,
	})

	return json, err
}
