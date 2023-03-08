package dto

type BankAccount struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Agency  string `json:"agency"`
	Account string `json:"account"`
	Pix     string `json:"pix"`
}
