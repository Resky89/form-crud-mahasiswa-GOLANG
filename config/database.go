package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection()(*sql.DB, error){
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "ipat08_uts"
	
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(127.0.0.1:3306)/"+dbName)
	return db, err

}

