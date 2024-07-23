package model

import (
	"database/sql"
)

type Model struct {
	Database *sql.DB
}
