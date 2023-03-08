package actions

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/builder"
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/dto"
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/repository"
)

type ClientActions struct {
	repository repository.ClientRepository
	builder    *builder.ClientBuilder
}

func NewClientAction(repository repository.ClientRepository) *ClientActions {
	return &ClientActions{
		repository: repository,
		builder:    builder.NewClientBuild(),
	}
}

func (c *ClientActions) CreateClient(clientDto dto.Client) dto.ActionOutput {

	client := c.builder.
		CreateClient(clientDto).
		WithAddress(clientDto.Addresses).
		WithBankAccounts(clientDto.BankAccounts).
		WithContacts(clientDto.Contacts).
		WithDocuments(clientDto.Documents).
		WithFiliation(clientDto.Filiation).
		Get()

	id, err := c.repository.Create(*client)

	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
		Data:    map[string]int{"id_created": id},
	}
}

func (c *ClientActions) UpdateClient(id int, clientDto dto.Client) dto.ActionOutput {
	client := c.builder.CreateClient(clientDto).Get()
	client.Id = id
	err := c.repository.Update(*client)

	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) DeleteClient(id int) dto.ActionOutput {
	c.repository.Delete(id)
	return dto.ActionOutput{
		Result:  true,
		Message: "Client deleted with success",
	}
}

func (c *ClientActions) CreateAddress(clientId int, addressDto dto.Address) dto.ActionOutput {
	address := c.builder.Address(addressDto)
	address.ClientId = clientId
	err := c.repository.CreateAddress(*address, nil)

	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) UpdateAddress(addressId int, addressDto dto.Address) dto.ActionOutput {
	address := c.builder.Address(addressDto)
	address.Id = addressId
	err := c.repository.UpdateAddress(*address)

	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) DeleteAddress(id int) dto.ActionOutput {
	c.repository.DeleteAddress(id)
	return dto.ActionOutput{
		Result:  true,
		Message: "Address deleted with success",
	}
}

func (c *ClientActions) CreateContact(clientId int, contactDto dto.Contact) dto.ActionOutput {
	contact := c.builder.Contact(contactDto)
	contact.ClientId = clientId
	err := c.repository.CreateContact(*contact, nil)

	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) UpdateContact(contactId int, contactDto dto.Contact) dto.ActionOutput {
	contact := c.builder.Contact(contactDto)
	contact.Id = contactId
	err := c.repository.UpdateContact(*contact)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) DeleteContact(contactId int) dto.ActionOutput {
	c.repository.DeleteContact(contactId)
	return dto.ActionOutput{
		Result:  true,
		Message: "Contact deleted with success",
	}
}

func (c *ClientActions) CreateBankAccount(clientId int, account dto.BankAccount) dto.ActionOutput {
	bAccount := c.builder.BankAccount(account)
	bAccount.ClientId = clientId
	err := c.repository.CreateBankAccount(*bAccount, nil)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) UpdateBankAccount(accountId int, accountDto dto.BankAccount) dto.ActionOutput {
	account := c.builder.BankAccount(accountDto)
	account.Id = accountId
	err := c.repository.UpdateBankAccount(*account)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) DeleteBankAccount(accountId int) dto.ActionOutput {
	c.repository.DeleteBankAccount(accountId)
	return dto.ActionOutput{
		Result:  true,
		Message: "Account deleted with success",
	}
}

func (c *ClientActions) CreateDocuments(clientId int, documentsDto dto.Documents) dto.ActionOutput {
	documents := c.builder.Documents(documentsDto)
	documents.ClientId = clientId
	err := c.repository.CreateDocuments(*documents, nil)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) UpdateDocuments(documentId int, documentDto dto.Documents) dto.ActionOutput {
	documents := c.builder.Documents(documentDto)
	documents.Id = documentId
	err := c.repository.UpdateDocuments(*documents)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) CreateFiliation(clientId int, filiatioDto dto.Filiation) dto.ActionOutput {
	filiation := c.builder.Filiation(filiatioDto)
	filiation.ClientId = clientId
	err := c.repository.CreateFiliation(*filiation, nil)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}

func (c *ClientActions) UpdateFiliation(filiationId int, documentDto dto.Filiation) dto.ActionOutput {
	filiation := c.builder.Filiation(documentDto)
	filiation.Id = filiationId
	err := c.repository.UpdateFiliation(*filiation)
	return dto.ActionOutput{
		Result:  err == nil,
		Message: err.Error(),
	}
}
