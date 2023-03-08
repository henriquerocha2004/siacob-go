package dto

type Documents struct {
	Rg            string `json:"rg"`
	CPForCNPJ     string `json:"cpf_or_cnpj"`
	TituloEleitor string `json:"titulo_eleitor"`
	CTPS          string `json:"ctps"`
	PIS           string `json:"pis"`
	CNH           string `json:"cnh"`
	Passport      string `json:"passport"`
	Reservista    string `json:"reservista"`
	IE            string `json:"ie"`
	IM            string `json:"im"`
}
