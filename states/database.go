package states

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB *sql.DB
)
