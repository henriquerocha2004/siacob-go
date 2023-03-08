package mysql

import (
	"time"

	"github.com/henriquerocha2004/siacob-go/internal/modules/clients/entities"
	"github.com/jmoiron/sqlx"
)

type ClientRepository struct {
	mysqlConnection *sqlx.DB
}

func NewClientRepository(con *sqlx.DB) *ClientRepository {
	return &ClientRepository{
		mysqlConnection: con,
	}
}

func (c *ClientRepository) Create(client entities.Client) (int, error) {
	transaction := c.mysqlConnection.MustBegin()
	client.SetDefaultTimeStamp()
	query := `
		INSERT INTO clients 
		(full_name, gender, type, birth_date, place_of_birth, nationality, marital_status, created_at, updated_at)
		VALUES
		(:full_name, :gender, :type, :birth_date, :place_of_birth, :nationality, :marital_status, :created_at, :updated_at)
	`
	result, err := transaction.NamedExec(query, client)
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	idInserted, err := result.LastInsertId()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	err = c.processAddress(int(idInserted), client.GetAddresses(), transaction)
	if err != nil {
		return 0, err
	}

	err = c.processBankAccounts(int(idInserted), client.GetBankAccounts(), transaction)
	if err != nil {
		return 0, err
	}

	err = c.processContacts(int(idInserted), client.GetContacts(), transaction)
	if err != nil {
		return 0, err
	}

	err = c.processDocuments(int(idInserted), client.GetDocuments(), transaction)
	if err != nil {
		return 0, err
	}

	err = c.processFiliation(int(idInserted), client.GetFiliation(), transaction)
	if err != nil {
		return 0, nil
	}

	err = transaction.Commit()
	if err != nil {
		transaction.Rollback()
		return 0, err
	}

	return int(idInserted), nil
}

func (c *ClientRepository) Update(client entities.Client) error {
	client.SetUpdatedAt()

	query := `
		UPDATE clients SET 
			full_name = :full_name, gender = :gender, type = :type, birth_date = :birth_date, 
			place_of_birth = :place_of_birth, nationality = :nationality, 
			marital_status = :marital_status, updated_at = :updated_at
		WHERE id = :id	  
	`
	_, err := c.mysqlConnection.NamedExec(query, client)
	return err
}

func (c *ClientRepository) Delete(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE clients SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}

func (c *ClientRepository) FindOne(id int) (*entities.Client, error) {
	client := entities.Client{}
	query := `SELECT * FROM clients WHERE id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Get(&client, query, id)
	if err != nil {
		return nil, err
	}

	addresses, err := c.FindAddressByClient(id)
	if err != nil {
		return &client, err
	}

	for _, address := range *addresses {
		client.AddAddress(address)
	}

	contacts, err := c.FindContactsByClient(id)
	if err != nil {
		return &client, err
	}

	for _, contact := range *contacts {
		client.AddContact(contact)
	}

	bankAccounts, err := c.FindBankAccountsByClient(id)
	if err != nil {
		return &client, err
	}

	for _, account := range *bankAccounts {
		client.AddBankAccount(account)
	}

	documents, err := c.FindDocumentsByClient(id)
	if err != nil {
		return &client, err
	}

	client.AddDocuments(documents)

	filiation, err := c.FindFiliationByClient(id)
	if err != nil {
		return &client, err
	}

	client.AddFiliation(filiation)

	return &client, err
}

func (c *ClientRepository) FindAll() ([]entities.Client, error) {
	clients := []entities.Client{}

	query := `SELECT * FROM clients`
	err := c.mysqlConnection.Select(&clients, query)
	return clients, err
}

func (c *ClientRepository) FindAddressByClient(clientId int) (*[]entities.Address, error) {
	addresses := []entities.Address{}
	query := `SELECT * FROM addresses WHERE client_id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Select(&addresses, query, clientId)
	if err != nil {
		return nil, err
	}

	return &addresses, nil
}

func (c *ClientRepository) FindContactsByClient(clientId int) (*[]entities.Contact, error) {
	contacts := []entities.Contact{}
	query := `SELECT * FROM contacts WHERE client_id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Select(&contacts, query, clientId)

	return &contacts, err
}

func (c *ClientRepository) FindBankAccountsByClient(clientId int) (*[]entities.BankAccount, error) {
	accounts := []entities.BankAccount{}
	query := `SELECT * FROM bank_accounts WHERE client_id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Select(&accounts, query, clientId)

	return &accounts, err
}

func (c *ClientRepository) FindDocumentsByClient(clientId int) (entities.Documents, error) {
	documents := entities.Documents{}
	query := `SELECT * FROM documents WHERE client_id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Get(&documents, query, clientId)

	return documents, err
}

func (c *ClientRepository) FindFiliationByClient(clientId int) (entities.Filiation, error) {
	filiation := entities.Filiation{}
	query := `SELECT * FROM filiation WHERE client_id = ? AND deleted_at IS NULL`
	err := c.mysqlConnection.Get(&filiation, query, clientId)

	return filiation, err
}

func (c *ClientRepository) processAddress(clientId int, addresses []entities.Address, transaction *sqlx.Tx) error {
	if len(addresses) < 1 {
		return nil
	}

	for _, address := range addresses {
		address.ClientId = clientId
		err := c.CreateAddress(address, transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientRepository) CreateAddress(address entities.Address, transaction *sqlx.Tx) error {

	if transaction == nil {
		transaction = c.mysqlConnection.MustBegin()
		defer transaction.Commit()
	}
	address.SetDefaultTimeStamp()
	query := `
		INSERT INTO addresses
		(street, district, city, zip_code, state, type, client_id, created_at, updated_at)
		VALUES
		(:street, :district, :city, :zip_code, :state, :type, :client_id, :created_at, :updated_at)
	`
	_, err := transaction.NamedExec(query, address)
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func (c *ClientRepository) UpdateAddress(address entities.Address) error {
	address.SetUpdatedAt()
	query := `
		UPDATE addresses SET 
			street = :street, district = :district,
			city = :city, zip_code = :zip_code, state = :state,
			type = :type, client_id = :client_id, updated_at = :updated_at 
		WHERE id = :id
	 `
	_, err := c.mysqlConnection.NamedExec(query, address)
	return err
}

func (c *ClientRepository) DeleteAddress(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE addresses SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}

func (c *ClientRepository) processBankAccounts(clientId int, accounts []entities.BankAccount, transaction *sqlx.Tx) error {
	if len(accounts) < 1 {
		return nil
	}

	for _, account := range accounts {
		account.ClientId = clientId
		err := c.CreateBankAccount(account, transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientRepository) CreateBankAccount(account entities.BankAccount, transaction *sqlx.Tx) error {
	if transaction == nil {
		transaction = c.mysqlConnection.MustBegin()
		defer transaction.Commit()
	}

	account.SetDefaultTimeStamp()

	query := `
		INSERT INTO bank_accounts
		(type, name, agency, account, pix, client_id, created_at, updated_at)
		VALUES
		(:type, :name, :agency, :account, :pix, :client_id, :created_at, :updated_at)
	`

	_, err := transaction.NamedExec(query, account)
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func (c *ClientRepository) UpdateBankAccount(account entities.BankAccount) error {
	account.SetUpdatedAt()
	query := `
		UPDATE bank_accounts SET 
			type = :type, 
			name = :name,
			agency = :agency, 
			account = :account, 
			pix = :pix,
			client_id = :client_id,
			updated_at = :updated_at 
		WHERE id = :id
	 `
	_, err := c.mysqlConnection.NamedExec(query, account)
	return err
}

func (c *ClientRepository) DeleteBankAccount(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE bank_accounts SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}

func (c *ClientRepository) processContacts(clientId int, contacts []entities.Contact, transaction *sqlx.Tx) error {
	if len(contacts) < 1 {
		return nil
	}

	for _, contact := range contacts {
		contact.ClientId = clientId
		err := c.CreateContact(contact, transaction)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *ClientRepository) CreateContact(contact entities.Contact, transaction *sqlx.Tx) error {
	if transaction == nil {
		transaction = c.mysqlConnection.MustBegin()
		defer transaction.Commit()
	}
	contact.SetDefaultTimeStamp()
	query := `
		INSERT INTO contacts
		(phone, mobile_operator, site, type, client_id, created_at, updated_at)
		VALUES
		(:phone, :mobile_operator, :site, :type, :client_id, :created_at, :updated_at)
	`

	_, err := transaction.NamedExec(query, contact)
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func (c *ClientRepository) UpdateContact(contact entities.Contact) error {
	contact.SetUpdatedAt()
	query := `
		UPDATE contacts SET 
			phone = :phone, 
			mobile_operator = :mobile_operator,
			site = :site, 
			type = :type, 
			client_id = :client_id,
			updated_at = :updated_at
		WHERE id = :id
	 `
	_, err := c.mysqlConnection.NamedExec(query, contact)
	return err
}

func (c *ClientRepository) DeleteContact(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE contacts SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}

func (c *ClientRepository) processDocuments(clientId int, documents *entities.Documents, transaction *sqlx.Tx) error {
	if documents == nil {
		return nil
	}
	documents.ClientId = clientId
	err := c.CreateDocuments(*documents, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientRepository) CreateDocuments(documents entities.Documents, transaction *sqlx.Tx) error {
	if transaction == nil {
		transaction = c.mysqlConnection.MustBegin()
		defer transaction.Commit()
	}

	documents.SetDefaultTimeStamp()

	query := `
		INSERT INTO documents
		(rg, cpf_or_cnpj, titulo_eleitor, ctps, pis, cnh, passport, reservista, ie, im, client_id, created_at, updated_at)
		VALUES
		(:rg, :cpf_or_cnpj, :titulo_eleitor, :ctps, :pis, :cnh, :passport, :reservista, :ie, :im, :client_id, :created_at, :updated_at)
	`
	_, err := transaction.NamedExec(query, documents)
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func (c *ClientRepository) UpdateDocuments(documents entities.Documents) error {
	documents.SetUpdatedAt()
	query := `
		UPDATE documents SET 
			rg = :rg, 
			cpf_or_cnpj = :cpf_or_cnpj,
			titulo_eleitor = :titulo_eleitor, 
			ctps = :ctps,
			pis = :pis,
			cnh = :cnh,
			passport = :passport,
			reservista = :reservista,
			ie = :ie,
			im = :im,  
			client_id = :client_id,
			updated_at = :updated_at
		WHERE id = :id
	 `
	_, err := c.mysqlConnection.NamedExec(query, documents)
	return err
}

func (c *ClientRepository) DeleteDocuments(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE documents SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}

func (c *ClientRepository) processFiliation(clientId int, filiation *entities.Filiation, transaction *sqlx.Tx) error {
	if filiation == nil {
		return nil
	}

	filiation.ClientId = clientId
	err := c.CreateFiliation(*filiation, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (c *ClientRepository) CreateFiliation(filiation entities.Filiation, transaction *sqlx.Tx) error {
	if transaction == nil {
		transaction = c.mysqlConnection.MustBegin()
		defer transaction.Commit()
	}

	filiation.SetDefaultTimeStamp()

	query := `
		INSERT INTO filiation
		(mother_name, father_name, client_id, created_at, updated_at)
		VALUES
		(:mother_name, :father_name, :client_id, created_at, updated_at)	
	`

	_, err := transaction.NamedExec(query, filiation)
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}

func (c *ClientRepository) UpdateFiliation(filiation entities.Filiation) error {
	filiation.SetUpdatedAt()
	query := `
		UPDATE filiation SET 
			mother_name = :mother_name, 
			father_name = :father_name, 
			client_id = :client_id,
			updated_at = :updated_at
		WHERE id = :id
	 `
	_, err := c.mysqlConnection.NamedExec(query, filiation)
	return err
}

func (c *ClientRepository) DeleteFiliation(id int) {
	deletedAt := time.Now().Format("2006-01-02 15:04")
	query := `UPDATE filiation SET deleted_at = ? WHERE id = ?`
	c.mysqlConnection.MustExec(query, deletedAt, id)
}
