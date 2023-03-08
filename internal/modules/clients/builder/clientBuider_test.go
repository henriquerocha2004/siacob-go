package builder

import (
	"testing"

	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/dto"
	"github.com/stretchr/testify/assert"
)

func TestBuildClient(t *testing.T) {
	t.Run("Should create client without additional information", func(t *testing.T) {
		clientDto := getDto()
		builder := NewClientBuild()
		client := builder.CreateClient(clientDto).Get()

		assert.Equal(t, clientDto.FullName, client.FullName)
		assert.Equal(t, clientDto.Gender, client.Gender.String)
		assert.Equal(t, clientDto.MaritalStatus, client.MaritalStatus.String)
		assert.Equal(t, clientDto.Nationality, client.Nationality.String)
		assert.Equal(t, clientDto.PlaceOfBirth, client.PlaceOfBirth.String)
		assert.Equal(t, clientDto.BirthDate, client.BirthDate.String)
		assert.Equal(t, clientDto.Type, client.Type.String)
	})
	t.Run("Should create client with address information", func(t *testing.T) {
		clientDto := getDto()
		builder := NewClientBuild()
		client := builder.CreateClient(clientDto).WithAddress(clientDto.Addresses).Get()

		assert.Equal(t, clientDto.Addresses[0].Street, client.GetAddresses()[0].Street.String)
		assert.Equal(t, clientDto.Addresses[0].District, client.GetAddresses()[0].District.String)
		assert.Equal(t, clientDto.Addresses[0].City, client.GetAddresses()[0].City.String)
		assert.Equal(t, clientDto.Addresses[0].State, client.GetAddresses()[0].State.String)
		assert.Equal(t, clientDto.Addresses[0].Type, client.GetAddresses()[0].Type.String)
		assert.Equal(t, clientDto.Addresses[0].ZipCode, client.GetAddresses()[0].ZipCode.String)
	})

	t.Run("Should create client with filiation information", func(t *testing.T) {
		clientDto := getDto()
		builder := NewClientBuild()
		client := builder.CreateClient(clientDto).WithFiliation(clientDto.Filiation).Get()

		assert.Equal(t, clientDto.Filiation.FatherName, client.GetFiliation().FatherName.String)
		assert.Equal(t, clientDto.Filiation.MotherName, client.GetFiliation().MotherName.String)
	})

	t.Run("Should create client with bank account information", func(t *testing.T) {
		clientDto := getDto()
		builder := NewClientBuild()
		client := builder.CreateClient(clientDto).WithBankAccounts(clientDto.BankAccounts).Get()

		assert.Equal(t, clientDto.BankAccounts[0].Account, client.GetBankAccounts()[0].Account.String)
		assert.Equal(t, clientDto.BankAccounts[0].Agency, client.GetBankAccounts()[0].Agency.String)
		assert.Equal(t, clientDto.BankAccounts[0].Name, client.GetBankAccounts()[0].Name.String)
		assert.Equal(t, clientDto.BankAccounts[0].Pix, client.GetBankAccounts()[0].Pix.String)
		assert.Equal(t, clientDto.BankAccounts[0].Type, client.GetBankAccounts()[0].Type.String)
	})

	t.Run("Should create client with documents information", func(t *testing.T) {
		clientDto := getDto()
		builder := NewClientBuild()
		client := builder.CreateClient(clientDto).WithDocuments(clientDto.Documents).Get()

		assert.Equal(t, clientDto.Documents.CNH, client.GetDocuments().CNH.String)
		assert.Equal(t, clientDto.Documents.CPForCNPJ, client.GetDocuments().CPForCNPJ.String)
		assert.Equal(t, clientDto.Documents.IE, client.GetDocuments().IE.String)
	})
}

func getDto() dto.Client {
	dtoClient := dto.Client{}
	dtoClient.FullName = "Henrique Rocha"
	dtoClient.Gender = "M"
	dtoClient.Type = "F"
	dtoClient.MaritalStatus = "S"
	dtoClient.BirthDate = "1999-01-01"
	dtoClient.Nationality = "Brasileiro"
	dtoClient.PlaceOfBirth = "Salvador"

	contact := dto.Contact{
		Phone:          "71956565656",
		MobileOperator: "VIVO",
		Site:           "www.algumsite.com",
		Type:           "P",
	}
	dtoClient.Contacts = append(dtoClient.Contacts, contact)

	address := dto.Address{
		Street:   "Rua dos Bobos N 0",
		District: "São Cristóvão",
		City:     "Salvador",
		State:    "BA",
		ZipCode:  "4157855",
		Type:     "P",
	}
	dtoClient.Addresses = append(dtoClient.Addresses, address)

	bankAccount := dto.BankAccount{
		Type:    "C",
		Name:    "Banco Jamelao",
		Agency:  "0012555454",
		Account: "5654546",
		Pix:     "56454654654564",
	}
	dtoClient.BankAccounts = append(dtoClient.BankAccounts, bankAccount)

	dtoClient.Documents = dto.Documents{
		Rg:            "123456789",
		CPForCNPJ:     "03547896544",
		TituloEleitor: "3125469781",
		CTPS:          "3265465456",
		PIS:           "5646546879879",
		CNH:           "6545466546",
		Passport:      "5645645456",
		Reservista:    "545456564",
		IE:            "545545445456",
		IM:            "5455454554",
	}

	dtoClient.Filiation = dto.Filiation{
		MotherName: "Maria dos Santos",
		FatherName: "José de Jesus",
	}

	return dtoClient
}
