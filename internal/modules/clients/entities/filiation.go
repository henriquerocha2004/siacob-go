package entities

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type Filiation struct {
	Id         int         `db:"id"`
	MotherName null.String `db:"mother_name" json:"mother_name"`
	FatherName null.String `db:"father_name" json:"father_name"`
	ClientId   int         `db:"client_id" json:"client_id"`
	timestamps.TimeStamps
}
