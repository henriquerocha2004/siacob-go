package dto

type Client struct {
	FullName      string        `json:"full_name"`
	Gender        string        `json:"gender"`
	Type          string        `json:"type"`
	BirthDate     string        `json:"birth_date"`
	PlaceOfBirth  string        `json:"place_of_birth"`
	Nationality   string        `json:"nationality"`
	MaritalStatus string        `json:"marital_status"`
	Addresses     []Address     `json:"addresses"`
	Documents     Documents     `json:"documents"`
	Contacts      []Contact     `json:"contacts"`
	BankAccounts  []BankAccount `json:"back_accounts"`
	Filiation     Filiation     `json:"filiation"`
}
