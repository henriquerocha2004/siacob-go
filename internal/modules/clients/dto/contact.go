package dto

type Contact struct {
	Phone          string `json:"phone"`
	MobileOperator string `json:"mobile_operator"`
	Site           string `json:"site"`
	Type           string `json:"type"`
}
