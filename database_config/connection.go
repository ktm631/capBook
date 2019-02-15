package database_config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DbConn() (db *gorm.DB) {
	dbDriver := "mysql"
	dbUser := DBUSER
	dbPass := DBPASSWORD
	dbHost := DBHOST
	dbPort := DBPORT
	dbName := DBNAME
	db, err := gorm.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
