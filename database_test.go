package go_mysql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
}

func TestOpenConnection(t *testing.T) {
	// Access MySQL on terminal:
	// /Applications/MAMP/Library/bin/mysql -uroot -proot
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
