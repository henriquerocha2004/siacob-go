package entities

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type BankAccount struct {
	Id       int         `db:"id"`
	Type     null.String `db:"type" json:"type"`
	Name     null.String `db:"name" json:"name"`
	Agency   null.String `db:"agency" json:"agency"`
	Account  null.String `db:"account" json:"account"`
	Pix      null.String `db:"pix" json:"pix"`
	ClientId int         `db:"client_id" json:"client_id"`
	timestamps.TimeStamps
}
