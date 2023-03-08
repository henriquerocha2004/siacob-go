package dto

type Address struct {
	Street   string `json:"street"`
	District string `json:"district"`
	City     string `json:"city"`
	ZipCode  string `json:"zip_code"`
	State    string `json:"state"`
	Type     string `json:"type"`
}
