package mysql_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/henriquerocha2004/siacob-go/internal/infra/database/mysql"
	"github.com/henriquerocha2004/siacob-go/internal/infra/database/mysql/testTools"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"gopkg.in/guregu/null.v4"

	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/builder"
	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/dto"
	"github.com/spf13/viper"
)

func init() {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo")
	viper.SetConfigName("../../../../env")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error in read file configuration %w", err))
	}
}

type TestClientSuit struct {
	suite.Suite
	connection    *sqlx.DB
	testTools     *testTools.DatabaseOperations
	clientBuilder *builder.ClientBuilder
	clientDto     dto.Client
	repository    *mysql.ClientRepository
}

func newTestClientSuit(connection *sqlx.DB, testTools *testTools.DatabaseOperations) *TestClientSuit {
	return &TestClientSuit{
		connection: connection,
		testTools:  testTools,
	}
}

func (s *TestClientSuit) SetupSuite() {
	s.clientBuilder = builder.NewClientBuild()
	s.clientDto = s.getDto()
	s.repository = mysql.NewClientRepository(s.connection)
}

func (s *TestClientSuit) BeforeTest(suiteName, testName string) {
	s.testTools.RefreshDatabase()
}

func TestManagerClient(t *testing.T) {
	connection := mysql.NewMysqlConnection()
	suite.Run(t, newTestClientSuit(connection, testTools.NewTestDatabaseOperations(connection)))
}

func (s *TestClientSuit) TestCreateClient() {
	s.Run("should create client without additional information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(id)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().Equal(s.clientDto.FullName, cliDb.FullName)
		s.Assert().Empty(cliDb.GetAddresses())
		s.Assert().Empty(cliDb.GetBankAccounts())
		s.Assert().False(cliDb.GetDocuments().CPForCNPJ.Valid)
		s.Assert().Empty(cliDb.GetContacts())
		s.Assert().False(cliDb.GetFiliation().FatherName.Valid)
	})
	s.Run("Should create client with address information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithAddress(s.clientDto.Addresses)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cliDb.GetAddresses())
	})
	s.Run("Should create client with contact information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithContacts(s.clientDto.Contacts).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cliDb.GetContacts())
	})
	s.Run("Should create client with bank accounts information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithBankAccounts(s.clientDto.BankAccounts).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cliDb.GetBankAccounts())
	})
	s.Run("Should create client with documents information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithDocuments(s.clientDto.Documents).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cliDb.GetDocuments())
	})
	s.Run("Should create client with filiation information", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithFiliation(s.clientDto.Filiation).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().NotEmpty(cliDb.GetFiliation())
	})
	s.Run("Should delete client", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		s.repository.Delete(id)
		_, err = s.repository.FindOne(id)
		s.Assert().Error(err)
		s.Assert().Equal("sql: no rows in result set", err.Error())
	})
	s.Run("should update client", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).Get()
		id, err := s.repository.Create(*client)
		s.Assert().NoError(err)
		client.FullName = "Carlos Antonio"
		client.Id = id
		err = s.repository.Update(*client)
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().Equal(client.FullName, cliDb.FullName)
	})
	s.Run("should update address", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithAddress(s.clientDto.Addresses)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		address := cliDb.GetAddresses()[0]
		address.Street = null.NewString("Rua da vitória", true)
		err = s.repository.UpdateAddress(address)
		s.Assert().NoError(err)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		addressDb := cliDb.GetAddresses()[0]
		s.Assert().Equal("Rua da vitória", addressDb.Street.String)
	})
	s.Run("Should delete address", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithAddress(s.clientDto.Addresses)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		address := cliDb.GetAddresses()[0]
		s.repository.DeleteAddress(address.Id)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().Empty(cliDb.GetAddresses())
	})
	s.Run("should update contact", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithContacts(s.clientDto.Contacts)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		contact := cliDb.GetContacts()[0]
		contact.Phone = null.NewString("719665656665", true)
		err = s.repository.UpdateContact(contact)
		s.Assert().NoError(err)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		contactDb := cliDb.GetContacts()[0]
		s.Assert().Equal("719665656665", contactDb.Phone.String)
	})
	s.Run("Should delete contact", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithContacts(s.clientDto.Contacts)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		contact := cliDb.GetContacts()[0]
		s.repository.DeleteContact(contact.Id)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().Empty(cliDb.GetContacts())
	})
	s.Run("should update bankAccount", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithBankAccounts(s.clientDto.BankAccounts)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		account := cliDb.GetBankAccounts()[0]
		account.Agency = null.NewString("71454545", true)
		err = s.repository.UpdateBankAccount(account)
		s.Assert().NoError(err)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		accountDb := cliDb.GetBankAccounts()[0]
		s.Assert().Equal("71454545", accountDb.Agency.String)
	})
	s.Run("Should delete bank Account", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithBankAccounts(s.clientDto.BankAccounts)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		account := cliDb.GetBankAccounts()[0]
		s.repository.DeleteBankAccount(account.Id)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		s.Assert().Empty(cliDb.GetBankAccounts())
	})
	s.Run("should update documents", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithDocuments(s.clientDto.Documents)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		documents := cliDb.GetDocuments()
		documents.Rg = null.NewString("789966666665", true)
		err = s.repository.UpdateDocuments(*documents)
		s.Assert().NoError(err)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		documentsDb := cliDb.GetDocuments()
		s.Assert().Equal("789966666665", documentsDb.Rg.String)
	})
	s.Run("should update filiation", func() {
		client := s.clientBuilder.CreateClient(s.clientDto).WithFiliation(s.clientDto.Filiation)
		id, err := s.repository.Create(*client.Get())
		s.Assert().NoError(err)
		cliDb, err := s.repository.FindOne(id)
		s.Assert().NoError(err)
		filiation := cliDb.GetFiliation()
		filiation.MotherName = null.NewString("Mamame", true)
		err = s.repository.UpdateFiliation(*filiation)
		s.Assert().NoError(err)
		cliDb, err = s.repository.FindOne(id)
		s.Assert().NoError(err)
		filiationDb := cliDb.GetFiliation()
		s.Assert().Equal("Mamame", filiationDb.MotherName.String)
	})
}

func (t *TestClientSuit) getDto() dto.Client {
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
