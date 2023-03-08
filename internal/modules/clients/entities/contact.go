package entities

import (
	"github.com/henriquerocha2004/siacob-go/internal/modules/helpers/timestamps"
	"gopkg.in/guregu/null.v4"
)

type Contact struct {
	Id             int         `db:"id"`
	Phone          null.String `db:"phone" json:"phone"`
	MobileOperator null.String `db:"mobile_operator" json:"mobile_operator"`
	Site           null.String `db:"site" json:"site"`
	Type           null.String `db:"type" json:"type"`
	ClientId       int         `db:"client_id" json:"client_id"`
	timestamps.TimeStamps
}
