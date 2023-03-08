package entities

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type Documents struct {
	Id            int         `db:"id"`
	Rg            null.String `db:"rg" json:"rg"`
	CPForCNPJ     null.String `db:"cpf_or_cnpj" json:"cpf_or_cnpj"`
	TituloEleitor null.String `db:"titulo_eleitor" json:"titulo_eleitor"`
	CTPS          null.String `db:"ctps" json:"ctps"`
	PIS           null.String `db:"pis" json:"pis"`
	CNH           null.String `db:"cnh" json:"cnh"`
	Passport      null.String `db:"passport" json:"passport"`
	Reservista    null.String `db:"reservista" json:"reservista"`
	IE            null.String `db:"ie" json:"ie"`
	IM            null.String `db:"im" json:"im"`
	ClientId      int         `db:"client_id" json:"client_id"`
	timestamps.TimeStamps
}
