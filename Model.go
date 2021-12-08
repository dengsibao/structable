package structable

import (
	"database/sql"
	"time"
)

type Model struct {
	ID        uint         `stbl:"id,PRIMARY_KEY,AUTO_INCREMENT"`
	CreatedAt time.Time    `stbl:"created_at"`
	UpdatedAt time.Time    `stbl:"updated_at"`
	DeletedAt sql.NullTime `stbl:"deleted_at"`
}
