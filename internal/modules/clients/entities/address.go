package entities

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type Address struct {
	Id       int         `db:"id" json:"id"`
	Street   null.String `db:"street" json:"street"`
	District null.String `db:"district" json:"district"`
	City     null.String `db:"city" json:"city"`
	ZipCode  null.String `db:"zip_code" json:"zip_code"`
	State    null.String `db:"state" json:"state"`
	Type     null.String `db:"type" json:"type"`
	ClientId int         `db:"client_id" json:"client_id"`
	timestamps.TimeStamps
}
