package timestamps

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type TimeStamps struct {
	CreatedAt string      `db:"created_at,omitempty"`
	UpdatedAt string      `db:"updated_at,omitempty"`
	DeletedAt null.String `db:"deleted_at,omitempty"`
}

func (t *TimeStamps) SetDefaultTimeStamp() {
	t.CreatedAt = time.Now().Format("2006-01-02 15:04")
	t.UpdatedAt = time.Now().Format("2006-01-02 15:04")
}

func (t *TimeStamps) SetUpdatedAt() {
	t.UpdatedAt = time.Now().Format("2006-01-02 15:04")
}

func (t *TimeStamps) SetDeletedAt() {
	t.DeletedAt = null.NewString(time.Now().Format("2006-01-02 15:04"), true)
}
